package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"database/sql"
	"fmt"
	golang_ture "golang-ture"
	"golang-ture/ent"
	"golang-ture/internal/handler"
	"golang-ture/internal/repositories"
	"golang-ture/internal/services"

	_ "golang-ture/docs"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//go:generate echo file $GOFILE is a part of $GOPACKAGE
// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8787
// @BasePath /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}
	client := OpenDB()

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		logrus.Fatal(err)
	}

	repository := repositories.NewRepository(client)
	service := services.NewService(repository)
	handlers := handler.NewHandler(service)
	if err := golang_ture.LoadTemplates(); err != nil {
		logrus.Fatalf("error detected while load templates: %s", err.Error())
	}

	srv := new(golang_ture.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error detected while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// Open new connection
func OpenDB() *ent.Client {
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"),
	)
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		logrus.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
