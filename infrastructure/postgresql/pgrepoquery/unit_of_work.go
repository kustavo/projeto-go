package pgrepoquery

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kustavo/benchmark/go/application/interfaces/repoquery"
)

type UnitOfWorkFactory struct {
	conn *pgxpool.Pool
}

func NewUnitOfWorkFactory(conn *pgxpool.Pool) repoquery.UnitOfWorkFactory {
	return &UnitOfWorkFactory{conn: conn}
}

func (uowfc *UnitOfWorkFactory) Create(ctx context.Context) repoquery.UnitOfWork {
	return newUnitOfWork(ctx, uowfc.conn)
}

type unitOfWork struct {
	ctx  context.Context
	conn *pgxpool.Pool
	tx   pgx.Tx
}

func newUnitOfWork(ctx context.Context, conn *pgxpool.Pool) *unitOfWork {
	return &unitOfWork{ctx: ctx, conn: conn}
}
