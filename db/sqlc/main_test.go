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
	dbSource = "postgresql://auth:ernkawezcvnwebj897472@db:5431/auth?sslmode=disable"
)

var testQueries *db.Queries
var dbtx db.DBTX

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot conect to db:", err)
	}

	testQueries = db.New()
	dbtx = conn

	os.Exit(m.Run())
}
