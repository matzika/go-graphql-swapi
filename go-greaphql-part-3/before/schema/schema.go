package schema

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-3/before/resolvers"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-3/before/types"
)

var (
	Schema graphql.Schema
)

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// 2
			"hello": &graphql.Field{
				Type: graphql.String,
				// 3
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", nil
				},
			},
			"human": &graphql.Field{
				Type: types.HumanType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the human",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}
					return resolvers.GetHuman(id), nil
				},
			},
			"humans": &graphql.Field{
				Type: graphql.NewList(types.HumanType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolvers.GetHumans(), nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: Query,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}
