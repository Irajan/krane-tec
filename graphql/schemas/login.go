package schema

import (
	"github.com/Irajan/krane-stack/resolvers"

	"github.com/graphql-go/graphql"
)

var LoginFields = graphql.Fields{
	"login": &graphql.Field{
		Type: resolvers.LoginType,
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: resolvers.Login,
	},
}
