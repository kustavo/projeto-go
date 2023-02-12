package pgrepoquery

import (
	"github.com/jackc/pgx/v4"
	"github.com/kustavo/benchmark/go/application/interfaces/repoquery"
	"github.com/kustavo/benchmark/go/domain/model"
)

type userRepo struct {
}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (userRepo *userRepo) GetByEmail(uow repoquery.UnitOfWork, email string) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE email=$1 FETCH FIRST 1 ROWS ONLY;"

	user := model.User{}
	row := QueryRow(uow, query, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepo) GetById(uow repoquery.UnitOfWork, id uint64) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id=$1;"

	user := model.User{}
	row := QueryRow(uow, query, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepo) GetUserCredentials(uow repoquery.UnitOfWork, username string) (*model.User, error) {
	query := "SELECT id, username, password FROM users WHERE username=$1;"

	user := model.User{}
	row := QueryRow(uow, query, username)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
