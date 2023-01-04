package repository_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/ThicoMoura/Auth/db/repository"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/jackc/pgx/v4/pgxpool"
)

var repo *repository.Repository

func TestMain(m *testing.M) {
	env, err := util.NewEnv("../../")
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	conn, err := pgxpool.Connect(context.Background(), env.Source)
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	repo = repository.NewRepository(db.NewStore(conn))

	os.Exit(m.Run())
}
