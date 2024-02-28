package gql

import (
	"errors"

	"github.com/graphql-go/graphql"

	model "github.com/muhammadali07/service-grap-go-api/app/domain/models"
	repository "github.com/muhammadali07/service-grap-go-api/app/domain/repository" // Import the repository package
)

// Define PostType (option 1: custom type)
var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Query root
var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: PostType, // Use PostType
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, errors.New("invalid argument")
				}

				// Option 1: Access repository directly (if in the same package)
				// post, err := repository.GetByID(id) // Assuming GetByID exists in the repository package

				// Option 2: Get repository from context (assuming context is passed)
				repo, ok := p.Context.Value("repository").(repository.PostRepository)
				if !ok {
					return nil, errors.New("missing repository in context")
				}
				post, err := repo.GetByID(id)

				if err != nil {
					return nil, err
				}
				return ToGraphQLPost(post), nil
			},
		},
	},
})

// Mutation root
var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": &graphql.Field{
			Type: PostType, // Use PostType
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				title, ok := p.Args["title"].(string)
				if !ok {
					return nil, errors.New("invalid argument")
				}
				content, ok := p.Args["content"].(string)
				if !ok {
					return nil, errors.New("invalid argument")
				}

				// Option 2: Get repository from context (assuming context is passed)
				repo, ok := p.Context.Value("repository").(repository.PostRepository)
				if !ok {
					return nil, errors.New("missing repository in context")
				}
				post, err := repo.Create(model.Post{
					Title:   title,
					Content: content,
				})

				if err != nil {
					return nil, err
				}
				return ToGraphQLPost(post), nil
			},
		},
	},
})

// Schema (ensures a single creation)
var Schema = graphql.NewSchema(graphql.SchemaConfig{
	Query:    Query,
	Mutation: Mutation,
})
