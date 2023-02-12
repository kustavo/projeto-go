package repocommand

import "context"

type UnitOfWorkFactory interface {
	Create(ctx context.Context) UnitOfWork
}

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
}
