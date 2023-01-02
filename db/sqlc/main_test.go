package db_test

import (
	"context"
	"log"
	"os"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgres://auth:we4naoir905adfh98u34235@db:5432/auth?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot conect to db:", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
