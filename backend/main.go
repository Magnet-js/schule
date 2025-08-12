package main

import (
	"fmt"
	"schule/config"
	"schule/internal/db"
)

func main() {
	cfg := config.MustLoadConfig()

	fmt.Println(cfg.DatabaseURL)
	database := db.MustConnectAndSetup(cfg.DatabaseURL)
	defer db.Close(database)
}