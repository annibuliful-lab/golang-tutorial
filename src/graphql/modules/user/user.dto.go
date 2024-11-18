package user

import "github.com/graph-gophers/graphql-go"

type User struct {
	Id    graphql.ID
	Email string
	Name  *string
}

type CreateUserInput struct {
	Id    graphql.ID
	Email string
	Name  *string
}

type OnCreatedUserSubscriber struct {
	Stop   <-chan struct{}
	Events chan<- *User
}
