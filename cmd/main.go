package main

import (
	"flag"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"someproject/config"
	"someproject/internal/api"
	"someproject/internal/repository"
)

func main() {
	dev := flag.Bool("dev",
		false,
		"enable reading config from .env file instead of system env vars",
	)
	flag.Parse()

	if *dev {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("postgres", conf.DSN)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUsersRepository(db)
	application := api.NewApplication(repo)
	application.Start(conf.ListenAddress())
}
