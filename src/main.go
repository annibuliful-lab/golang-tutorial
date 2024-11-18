package main

import (
	"log"
	"net/http"
	"os"

	localGraphql "backend/src/graphql"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
)

func main() {

	graphqlSchema, err := os.ReadFile("./src/graphql/schema.graphql")

	if err != nil {
		panic(err)
	}

	s := string(graphqlSchema)
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(s, localGraphql.NewResolver(), opts...)
	graphQLHandler := graphqlws.NewHandlerFunc(schema, &relay.Handler{Schema: schema})
	http.Handle("/query", graphQLHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
