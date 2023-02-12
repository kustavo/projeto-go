package repocommand

import (
	"github.com/kustavo/benchmark/go/domain/model"
)

type UserRepo interface {
	//Fetch(num int64) ([]*models.Post, error)
	Create(uow UnitOfWork, u *model.User) (uint64, error)
	Update(uow UnitOfWork, u *model.User) error
	GetById(uow UnitOfWork, id uint64) (*model.User, error)
	Delete(uow UnitOfWork, id uint64) error
}
