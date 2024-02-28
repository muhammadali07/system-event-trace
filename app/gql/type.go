package gql

import model "github.com/muhammadali07/service-grap-go-api/app/domain/models"

// Tipe GraphQL untuk Post
type Post struct {
	ID      int    `graphql:"id"`
	Title   string `graphql:"title"`
	Content string `graphql:"content"`
}

// Konversi dari model Post ke GraphQL Post
func ToGraphQLPost(post model.Post) Post {
	return Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
}

// Konversi dari GraphQL Post ke model Post
func FromGraphQLPost(post Post) model.Post {
	return model.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
}

// Tipe GraphQL untuk User
type User struct {
	ID    int    `graphql:"id"`
	Name  string `graphql:"name"`
	Email string `graphql:"email"`
}

// Konversi dari model User ke GraphQL User
func ToGraphQLUser(user model.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

// Konversi dari GraphQL User ke model User
func FromGraphQLUser(user User) model.User {
	return model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
