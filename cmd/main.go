package main

import (
	"context"
	"github.com/chapdast/agileful"
	"github.com/chapdast/agileful/db"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	DB, err := db.NewDB(ctx, db.GetDBLinkFromEnv())
	if err != nil {
		log.Fatal(err)
	}
	app, err := agileful.Run(DB)
	if err != nil {
		log.Fatal(err)
	}
	if err = app.Listen(GetServingPort()); err != nil {
		log.Fatal(err)
	}
}

func GetServingPort() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return ":3000"
	}
	return port
}
