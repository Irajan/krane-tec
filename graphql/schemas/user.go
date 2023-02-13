package schema

import (
	"github.com/Irajan/krane-stack/resolvers"

	"github.com/graphql-go/graphql"
)

var RegisterFields = graphql.Fields{
	"register": &graphql.Field{
		Type: resolvers.UserType,
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"fullName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"role": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: resolvers.RegisterUser,
	},
}
