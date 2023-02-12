package pgrepocommand

import (
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/kustavo/benchmark/go/application/interfaces/repocommand"
	"github.com/kustavo/benchmark/go/domain/model"
)

type userRepo struct {
}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (ur *userRepo) Create(uow repocommand.UnitOfWork, u *model.User) (uint64, error) {
	query := "INSERT INTO users (created_at, updated_at, username, password, name, email) VALUES ($1, $2, $3, $4, $5, $6) returning id;"

	var id uint64
	dateNow := time.Now()
	err := QueryRow(uow, query, dateNow, dateNow, &u.Username, &u.Password, &u.Name, &u.Email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *userRepo) Update(uow repocommand.UnitOfWork, u *model.User) error {
	query := "UPDATE users SET updated_at=$1, name=$2, email=$3 WHERE id=$4;"

	cmdtag, err := Exec(uow, query, time.Now(), &u.Name, &u.Email, &u.ID)
	if cmdtag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return err
}

func (ur *userRepo) GetById(uow repocommand.UnitOfWork, id uint64) (*model.User, error) {
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

func (ur *userRepo) Delete(uow repocommand.UnitOfWork, id uint64) error {
	query := "DELETE FROM users WHERE id=$1;"

	cmdtag, err := Exec(uow, query, id)
	if cmdtag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return err
}
