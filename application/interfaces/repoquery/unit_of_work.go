package repoquery

import "context"

type UnitOfWorkFactory interface {
	Create(ctx context.Context) UnitOfWork
}

type UnitOfWork interface {
}
