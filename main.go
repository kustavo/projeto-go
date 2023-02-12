package main

import (
	"context"
	"log"
	"net/http"
	"path"

	"github.com/kustavo/benchmark/go/domain/message"
	"github.com/kustavo/benchmark/go/infrastructure"
	"github.com/kustavo/benchmark/go/infrastructure/http/router"
	"github.com/kustavo/benchmark/go/infrastructure/postgresql"
	"github.com/kustavo/benchmark/go/infrastructure/redis"

	"github.com/joho/godotenv"
)

func main() {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panicln(message.ErrLoadingEnvFile, message.Separator, err)
	}

	// database connection
	conn := postgresql.CreateDataBaseConnection(context.Background())
	defer conn.Close()

	// redis connection
	redis, err := redis.NewRedis()
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	// application
	application := infrastructure.NewApplication(conn, conn, redis)
	r := router.NewRouter(application.Authentication)
	r.AddRoutes(router.GetUserRouter(application))

	// http.HandleFunc("/doc", docs)

	// http server
	log.Fatal(http.ListenAndServe(":3000", r))
}

func docs(w http.ResponseWriter, r *http.Request) {
	fp := path.Join(".", "swagger.json")
	http.ServeFile(w, r, fp)

	// opts := middleware.SwaggerUIOpts{SpecURL: "swagger.json"}
	// sh := middleware.SwaggerUI(opts, nil)
	// http.Handle("/docs", sh)
}
