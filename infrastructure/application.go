package infrastructure

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kustavo/benchmark/go/application/command"
	"github.com/kustavo/benchmark/go/application/interfaces"
	"github.com/kustavo/benchmark/go/application/query"
	"github.com/kustavo/benchmark/go/infrastructure/jwt"
	"github.com/kustavo/benchmark/go/infrastructure/postgresql/pgrepocommand"
	"github.com/kustavo/benchmark/go/infrastructure/postgresql/pgrepoquery"
)

func NewApplication(connCmd *pgxpool.Pool, connQy *pgxpool.Pool, cache interfaces.Cache) *interfaces.Application {

	uowFacCmd := pgrepocommand.NewUnitOfWorkFactory(connCmd)
	userRepoCmd := pgrepocommand.NewUserRepo()
	createUserCommand := command.NewCreateUserCommand(uowFacCmd, userRepoCmd)
	deleteUserCommand := command.NewDeleteUserCommand(uowFacCmd, userRepoCmd)
	updateUserCommand := command.NewUpdateUserCommand(uowFacCmd, userRepoCmd)

	uowFacQy := pgrepoquery.NewUnitOfWorkFactory(connQy)
	userRepoQy := pgrepoquery.NewUserRepo()
	getUserByEmailQuery := query.NewGetUserByEmailQuery(uowFacQy, userRepoQy)
	getUserByIdQuery := query.NewGetUserByIdQuery(uowFacQy, userRepoQy)
	getUserCredentialsQuery := query.NewGetUserCredentialsQuery(uowFacQy, userRepoQy)

	commands := interfaces.Commands{
		CreateUserCommand: createUserCommand,
		DeleteUserCommand: *deleteUserCommand,
		UpdateUserCommand: *updateUserCommand,
	}
	queries := interfaces.Queries{
		GetUserByEmailQuery:     *getUserByEmailQuery,
		GetUserByIdQuery:        *getUserByIdQuery,
		GetUserCredentialsQuery: *getUserCredentialsQuery,
	}

	auth := jwt.NewJwt(cache)

	return &interfaces.Application{Commands: commands, Queries: queries, Authentication: auth}
}
