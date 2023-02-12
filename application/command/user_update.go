package command

import (
	"context"

	"github.com/kustavo/benchmark/go/application/interfaces/repocommand"

	"github.com/kustavo/benchmark/go/domain/model"
)

type UpdateUserCommand struct {
	uowFactory repocommand.UnitOfWorkFactory
	userRepo   repocommand.UserRepo
}

func NewUpdateUserCommand(uowFactory repocommand.UnitOfWorkFactory, userRepo repocommand.UserRepo) *UpdateUserCommand {
	return &UpdateUserCommand{uowFactory: uowFactory, userRepo: userRepo}
}

func (cmd *UpdateUserCommand) Handle(ctx context.Context, u *model.User) error {

	uow := cmd.uowFactory.Create(ctx)

	err := uow.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			uow.Rollback()
		} else {
			uow.Commit()
		}
	}()

	err = cmd.userRepo.Update(uow, u)
	return err
}
