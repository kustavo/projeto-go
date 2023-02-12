package controller_test

import (
	"context"
	"testing"

	"github.com/kustavo/benchmark/go/application/command"
	"github.com/kustavo/benchmark/go/application/interfaces"
	"github.com/kustavo/benchmark/go/infrastructure/http/controller"
)

type CreateUserCommand struct{}

func (cmd *CreateUserCommand) Handle(ctx context.Context, u *command.CreateUserRequest) (uint64, error) {
	return 1, nil
}

func TestUserController_Create(t *testing.T) {

	cmd := CreateUserCommand{}

	commands := interfaces.Commands{
		CreateUserCommand: &cmd,
	}

	app := &interfaces.Application{
		Commands: commands,
	}

	ctrl := controller.NewUserController(app)

	_ = ctrl
}
