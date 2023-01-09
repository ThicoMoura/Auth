package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThicoMoura/Auth/api/controller"
	"github.com/ThicoMoura/Auth/db/repository"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/token"
	"github.com/ThicoMoura/Auth/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	loc, _ := time.LoadLocation("America/Campo_Grande")

	Lerror := util.NewLogs(os.Stderr, util.Lerror, loc)
	Linfo := util.NewLogs(os.Stderr, util.Linfo, loc)

	env, err := util.NewEnv("./")
	if err != nil {
		Lerror.Fatal(err)
	}

	m, err := migrate.New(env.Migrate, env.Source)
	if err != nil {
		Lerror.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		Lerror.Fatal(err)
	}

	token, err := token.NewPaseto(env.Key)
	if err != nil {
		Lerror.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conn, err := pgxpool.Connect(ctx, env.Source)
	if err != nil {
		Lerror.Fatal(err)
	}

	gin.SetMode(env.GinMode)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		switch httpMethod {
		case "POST":
			Linfo.Printf("endpoint: \"POST\"   ->  %v\n", absolutePath)
		case "GET":
			Linfo.Printf("endpoint: \"GET\"    ->  %v\n", absolutePath)
		case "PUT":
			Linfo.Printf("endpoint: \"PUT\"    ->  %v\n", absolutePath)
		case "DELETE":
			Linfo.Printf("endpoint: \"DELETE\" ->  %v\n", absolutePath)
		}

	}

	server := controller.NewServer(repository.NewRepository(db.NewStore(conn)), token)

	srv := server.Start("0.0.0.0:80")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Lerror.Fatal(err)
		}
	}()

	<-ctx.Done()

	stop()
	Linfo.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		Lerror.Fatal(err)
	}

	Linfo.Println("Server exiting")
}
