package main

import (
	"context"
	"log"

	"github.com/ThicoMoura/Auth/api/controller"
	"github.com/ThicoMoura/Auth/db/repository"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	env, err := util.NewEnv("./")
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	m, err := migrate.New(env.Migrate, env.Source)
	if err != nil {
		log.Fatal("cannot create a migrate instace: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("cannot migrate up: ", err)
	}

	conn, err := pgxpool.Connect(context.Background(), env.Source)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	gin.SetMode(env.GinMode)

	server := controller.NewServer(repository.NewRepository(db.NewStore(conn)))

	server.Start("0.0.0.0:80")
}
