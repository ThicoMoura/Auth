package db_test

import (
	"context"
	"log"
	"os"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	env, err := util.NewEnv("../../")
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	conn, err := pgx.Connect(context.Background(), env.Source)
	if err != nil {
		log.Fatal("cannot conect to db: ", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
