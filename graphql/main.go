package main

import (
	"github.com/Irajan/krane-stack/configs"
	"github.com/Irajan/krane-stack/db"
	schema "github.com/Irajan/krane-stack/schemas"
	errorUtils "github.com/Irajan/krane-stack/utils/error"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {
	configs.InitializeConfigs()
	db.Initialize()

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	})

	schema := schema.InitializeSchema()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r.POST("/graphql", gin.WrapH(h))
	err := r.Run()

	errorUtils.HandleError(err)

	defer db.Close()

}
