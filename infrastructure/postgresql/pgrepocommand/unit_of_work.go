package pgrepocommand

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kustavo/benchmark/go/application/interfaces/repocommand"
)

type UnitOfWorkFactory struct {
	conn *pgxpool.Pool
}

func NewUnitOfWorkFactory(conn *pgxpool.Pool) repocommand.UnitOfWorkFactory {
	return &UnitOfWorkFactory{conn: conn}
}

func (uowfc *UnitOfWorkFactory) Create(ctx context.Context) repocommand.UnitOfWork {
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

func (uow *unitOfWork) Begin() error {
	var err error
	uow.tx, err = uow.conn.BeginTx(uow.ctx, pgx.TxOptions{})
	return err
}

func (uow *unitOfWork) Commit() error {
	err := uow.tx.Commit(uow.ctx)
	uow.tx = nil
	return err
}

func (uow *unitOfWork) Rollback() error {
	err := uow.tx.Rollback(uow.ctx)
	uow.tx = nil
	return err
}
