package query

import (
	"context"

	"github.com/kustavo/benchmark/go/application/interfaces/repoquery"

	"github.com/kustavo/benchmark/go/domain/model"
)

type GetUserByEmailQuery struct {
	uowFactory repoquery.UnitOfWorkFactory
	userRepo   repoquery.UserRepo
}

func NewGetUserByEmailQuery(uowFactory repoquery.UnitOfWorkFactory, userRepo repoquery.UserRepo) *GetUserByEmailQuery {
	return &GetUserByEmailQuery{uowFactory: uowFactory, userRepo: userRepo}
}

func (qy *GetUserByEmailQuery) Handle(ctx context.Context, email string) (*model.User, error) {

	uow := qy.uowFactory.Create(ctx)

	user, err := qy.userRepo.GetByEmail(uow, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
