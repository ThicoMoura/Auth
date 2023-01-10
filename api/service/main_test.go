package service_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/ThicoMoura/Auth/api/service"
	"github.com/ThicoMoura/Auth/db/repository"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/jackc/pgx/v4/pgxpool"
)

var services map[string]service.Service

func TestMain(m *testing.M) {
	env, err := util.NewEnv("../../")
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	conn, err := pgxpool.Connect(context.Background(), env.Source)
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	repo := repository.NewRepository(db.NewStore(conn))

	services = map[string]service.Service{
		"access": service.NewAccess(repo.Table("access")),
		"group":  service.NewGroup(repo.Table("group")),
		"login": service.NewLogin(map[string]repository.IRepository{
			"user":    repo.Table("user"),
			"session": repo.Table("session"),
		}),
		"session": service.NewSession(repo.Table("session")),
		"system":  service.NewSystem(repo.Table("system")),
		"user":    service.NewUser(repo.Table("user")),
	}

	os.Exit(m.Run())
}
