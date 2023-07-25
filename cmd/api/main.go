package main

import (
	"log"

	"github.com/eduardylopes/link-shortener/configs/db"
	"github.com/eduardylopes/link-shortener/internal/code"
	"github.com/eduardylopes/link-shortener/internal/link"
	"github.com/eduardylopes/link-shortener/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	codeSvc := code.NewService()
	linkRep := link.NewRepository(dbConn.GetDB())
	linkSvc := link.NewService(linkRep, codeSvc)
	linkHandler := link.NewHandler(linkSvc)

	router.InitRouter(linkHandler)
	router.Start()
}
