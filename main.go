package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conn, err := pgxpool.Connect(ctx, env.Source)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	gin.SetMode(env.GinMode)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v\n", httpMethod, absolutePath)
	}

	server := controller.NewServer(repository.NewRepository(db.NewStore(conn)))

	srv := server.Start("0.0.0.0:80")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
