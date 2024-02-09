package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Proger30/exchangeRatesCBRF/internal/config"
	"github.com/Proger30/exchangeRatesCBRF/internal/database"
	http_calls "github.com/Proger30/exchangeRatesCBRF/internal/http-calls"
	"github.com/Proger30/exchangeRatesCBRF/internal/router"
)

var (
	db  *sqlx.DB
	cfg config.Config
)

func main() {
	var err error
	cfg, err = config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		err = nil
	}
	db, err = database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		err = nil
	}
	go updateCurrenciesPeriodically()

	r := router.SetupRouter(db)
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func updateCurrenciesPeriodically() {
	for {
		http_calls.UpdateCurrencies(db)
		time.Sleep(time.Duration(cfg.UpdateTime) * time.Second)
	}
}
