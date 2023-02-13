package resolvers

import (
	"errors"
	"time"

	errorUtils "github.com/Irajan/krane-stack/utils/error"

	"github.com/Irajan/krane-stack/db"

	"github.com/Irajan/krane-stack/configs"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	"github.com/graphql-go/graphql"
)

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	UserId       int64  `json:"userId"`
	RefreshToken string `json:"refreshToken"`
}

var LoginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthorizedUser",
		Fields: graphql.Fields{
			"accessToken": &graphql.Field{
				Type: graphql.String,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"refreshToken": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func Login(p graphql.ResolveParams) (interface{}, error) {

	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	rows := db.Query(`SELECT id, name as fullName, role, password_hash as passwordHash from users where email= $1 `, email)

	// If No rows found for given email so return error
	if !rows.Next() {
		return nil, errors.New("email not registered")
	}

	defer rows.Close()

	var passwordHash string
	var fullName string
	var role string
	var userId int64

	err := rows.Scan(&userId, &passwordHash, &fullName, &role)
	errorUtils.HandleError(err, "Error on scanning data")

	// If password is incorrect return error
	if checkPasswordHash(password, passwordHash) {
		return nil, errors.New("invalid credentials")
	}

	// Generate jwt tokens

	accessTokenClaims := jwt.MapClaims{
		"iss":      "issuer",
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
		"email":    email,
		"fullName": fullName,
		"role":     role,
	}
	refreshTokenClaims := jwt.MapClaims{
		"iss": "issuer",
		"exp": time.Now().Add(time.Hour * 2400).Unix(),
		"iat": time.Now().Unix(),
	}

	return loginResponse{
		AccessToken:  generateJwtTokens(accessTokenClaims, configs.Config.JwtAccessToken),
		RefreshToken: generateJwtTokens(refreshTokenClaims, configs.Config.JwtRefreshToken),
		UserId:       userId,
	}, nil

}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	errorUtils.HandleError(err)

	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJwtTokens(claims jwt.Claims, secretKey string) string {
	encodedRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := encodedRefreshToken.SignedString([]byte(secretKey))
	errorUtils.HandleError(err, "Error generating token")

	return token
}
