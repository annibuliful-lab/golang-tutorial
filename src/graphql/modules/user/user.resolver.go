package user

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type UserResolver struct{}

func (UserResolver) CreateUser(ctx context.Context, input struct {
	Input CreateUserInput
}) User {

	return User{
		Id:    input.Input.Id,
		Email: input.Input.Email,
		Name:  input.Input.Name,
	}
}

func (UserResolver) GetUsers() []User {

	var users []User

	users = append(users, User{
		Id:    graphql.ID("1"),
		Email: "test@test.com",
	})

	return users
}

func (UserResolver) GetUserById(ctx context.Context, input struct {
	Id graphql.ID
}) User {

	var name = string("test")

	return User{
		Id:    input.Id,
		Email: "test@test.com",
		Name:  &name,
	}
}

func (UserResolver) OnCreatedUser(ctx context.Context) User {
	return User{}
}
