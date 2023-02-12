package postgresql

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateDataBaseConnection(ctx context.Context) *pgxpool.Pool {
	conf, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL_PGX"))
	if err != nil {
		log.Fatalln("database parseconfig error:", err)
		os.Exit(1)
	}
	conf.LazyConnect = true

	conn, err := pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		log.Fatalln("unable to connect to database:", err)
		os.Exit(1)
	}

	return conn
}
