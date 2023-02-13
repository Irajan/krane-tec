package resolvers

import (
	"errors"

	"github.com/Irajan/krane-stack/db"

	"github.com/graphql-go/graphql"
)

type userResponse struct {
	Id       int64  `json:"id"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"fullName": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func RegisterUser(p graphql.ResolveParams) (interface{}, error) {

	email := p.Args["email"].(string)
	fullName := p.Args["fullName"].(string)
	role := p.Args["role"].(string)

	//Check if the email already exists or not

	rows := db.Query(`SELECT email from users where email = $1`, email)

	if rows.Next() {
		return nil, errors.New("email already registered")
	}

	defer rows.Close()

	// Hash the password
	password := p.Args["password"].(string)
	hashedPassword := hashPassword(password)

	var userId int64

	db.Insert(`INSERT INTO users 
						(name,password_hash,email,role)
					VALUES ($1,$2,$3,$4)
						returning id`, []any{fullName, hashedPassword, email, role}, []any{&userId})

	return userResponse{
		Id:       userId,
		FullName: fullName,
		Role:     role,
		Email:    email,
	}, nil
}
