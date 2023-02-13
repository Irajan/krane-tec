package schema

import (
	"fmt"

	logger "github.com/Irajan/krane-stack/utils/logger"

	"github.com/graphql-go/graphql"
)

func initializeMutation() *graphql.Object {

	fields := graphql.Fields(combineFields(LoginFields, RegisterFields))

	mutationObject := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: fields,
	})

	return mutationObject
}

func initializeQuery() *graphql.Object {

	return &graphql.Object{}

}

func InitializeSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    initializeQuery(),
			Mutation: initializeMutation(),
		},
	)

	if err != nil {
		logger.Error("Error initializing graph ql schema")

		panic(err)
	}

	return schema

}

func combineFields(fieldMaps ...map[string]*graphql.Field) map[string]*graphql.Field {
	result := make(map[string]*graphql.Field)
	for _, m := range fieldMaps {
		for k, v := range m {
			result[k] = v
		}
	}

	fmt.Printf("%T", result)

	return result
}
