package command

import (
	"context"

	"github.com/kustavo/benchmark/go/application/interfaces/repocommand"

	"github.com/kustavo/benchmark/go/domain/model"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type CreateUserCommandInterface interface {
	Handle(ctx context.Context, u *CreateUserRequest) (uint64, error)
}

type CreateUserCommand struct {
	uowFactory repocommand.UnitOfWorkFactory
	userRepo   repocommand.UserRepo
}

func NewCreateUserCommand(uowFactory repocommand.UnitOfWorkFactory, userRepo repocommand.UserRepo) *CreateUserCommand {
	return &CreateUserCommand{uowFactory: uowFactory, userRepo: userRepo}
}

func (cmd *CreateUserCommand) Handle(ctx context.Context, u *CreateUserRequest) (uint64, error) {

	user, err := model.NewUser(u.Username, u.Password, u.Name, u.Email)
	if err != nil {
		return 0, err
	}

	uow := cmd.uowFactory.Create(ctx)

	err = uow.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			uow.Rollback()
		} else {
			uow.Commit()
		}
	}()

	newID, err := cmd.userRepo.Create(uow, user)
	if err != nil {
		return 0, err
	}

	return newID, nil
}
