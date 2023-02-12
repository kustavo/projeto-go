package query

import (
	"context"

	"github.com/kustavo/benchmark/go/application/interfaces/repoquery"

	"github.com/kustavo/benchmark/go/domain/model"
)

type GetUserByIdQuery struct {
	uowFactory repoquery.UnitOfWorkFactory
	userRepo   repoquery.UserRepo
}

func NewGetUserByIdQuery(uowFactory repoquery.UnitOfWorkFactory, userRepo repoquery.UserRepo) *GetUserByIdQuery {
	return &GetUserByIdQuery{uowFactory: uowFactory, userRepo: userRepo}
}

func (qy *GetUserByIdQuery) Handle(ctx context.Context, id uint64) (*model.User, error) {

	uow := qy.uowFactory.Create(ctx)

	user, err := qy.userRepo.GetById(uow, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
