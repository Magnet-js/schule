package main

import (
	"schule/config"
	"schule/internal/db"
)

func main() {
	cfg := config.MustLoadConfig()

	database := db.MustConnectAndSetup(cfg.DatabaseURL)
	defer db.Close(database)
}