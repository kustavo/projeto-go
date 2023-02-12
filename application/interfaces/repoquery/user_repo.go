package repoquery

import (
	"github.com/kustavo/benchmark/go/domain/model"
)

type UserRepo interface {
	GetByEmail(uow UnitOfWork, email string) (*model.User, error)
	GetById(uow UnitOfWork, id uint64) (*model.User, error)
	GetUserCredentials(uow UnitOfWork, username string) (*model.User, error)
}
