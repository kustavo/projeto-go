package pgrepocommand

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/kustavo/benchmark/go/application/interfaces/repocommand"
)

func QueryRow(uow repocommand.UnitOfWork, query string, args ...interface{}) pgx.Row {
	uowPg := uow.(*unitOfWork)
	if uowPg.tx != nil {
		return uowPg.tx.QueryRow(uowPg.ctx, query, args...)
	} else {
		return uowPg.conn.QueryRow(uowPg.ctx, query, args...)
	}
}

func Exec(uow repocommand.UnitOfWork, query string, args ...interface{}) (pgconn.CommandTag, error) {
	uowPg := uow.(*unitOfWork)
	if uowPg.tx != nil {
		return uowPg.tx.Exec(uowPg.ctx, query, args...)
	} else {
		return uowPg.conn.Exec(uowPg.ctx, query, args...)
	}
}
