package query

import (
	"context"

	"github.com/kustavo/benchmark/go/application/interfaces/repoquery"

	"github.com/kustavo/benchmark/go/domain/message"
	"github.com/kustavo/benchmark/go/domain/model"
)

type GetUserCredentialsQuery struct {
	uowFactory repoquery.UnitOfWorkFactory
	userRepo   repoquery.UserRepo
}

func NewGetUserCredentialsQuery(uowFactory repoquery.UnitOfWorkFactory, userRepo repoquery.UserRepo) *GetUserCredentialsQuery {
	return &GetUserCredentialsQuery{uowFactory: uowFactory, userRepo: userRepo}
}

func (qy *GetUserCredentialsQuery) Handle(ctx context.Context, username string, password string) (*model.User, error) {

	uow := qy.uowFactory.Create(ctx)

	user, err := qy.userRepo.GetUserCredentials(uow, username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, message.ErrIncorrectUsernameOrPassword
	}

	return user, nil
}
