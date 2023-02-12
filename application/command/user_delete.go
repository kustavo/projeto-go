package command

import (
	"context"

	"github.com/kustavo/benchmark/go/application/interfaces/repocommand"
)

type DeleteUserCommand struct {
	uowFactory repocommand.UnitOfWorkFactory
	userRepo   repocommand.UserRepo
}

func NewDeleteUserCommand(uowFactory repocommand.UnitOfWorkFactory, userRepo repocommand.UserRepo) *DeleteUserCommand {
	return &DeleteUserCommand{uowFactory: uowFactory, userRepo: userRepo}
}

func (cmd *DeleteUserCommand) Handle(ctx context.Context, id uint64) error {

	uow := cmd.uowFactory.Create(ctx)

	err := cmd.userRepo.Delete(uow, id)
	return err
}
