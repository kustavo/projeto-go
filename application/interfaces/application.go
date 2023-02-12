package interfaces

import (
	"github.com/kustavo/benchmark/go/application/command"
	"github.com/kustavo/benchmark/go/application/query"
)

type Application struct {
	Commands       Commands
	Queries        Queries
	Authentication Authentication
}

type Commands struct {
	CreateUserCommand command.CreateUserCommandInterface
	CreateUserRequest command.CreateUserRequest
	DeleteUserCommand command.DeleteUserCommand
	UpdateUserCommand command.UpdateUserCommand
}

type Queries struct {
	GetUserByEmailQuery     query.GetUserByEmailQuery
	GetUserByIdQuery        query.GetUserByIdQuery
	GetUserCredentialsQuery query.GetUserCredentialsQuery
}
