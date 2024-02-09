package http_calls

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type Currency struct {
	Id      int       `json:"id" db:"id"`
	Code    string    `json:"code" db:"code"`
	Rate    float64   `json:"rate" db:"rate"`
	Updated time.Time `json:"updated" db:"updated"`
}

type CurrencyResponce struct {
	Disclaimer string
	Date       string
	Timestamp  int
	Base       string
	Rates      map[string]float64
}

func UpdateCurrencies(db *sqlx.DB) {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		log.Printf("Failed to fetch currency rates: %v", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return
	}

	updateCurrenciesFromJSON(data, db)
}

func updateCurrenciesFromJSON(data []byte, db *sqlx.DB) {
	var rates CurrencyResponce
	if err := json.Unmarshal(data, &rates); err != nil {
		log.Printf("Failed to decode JSON: %v", err)
		return
	}

	tx := db.MustBegin()
	defer tx.Rollback()

	for code, rate := range rates.Rates {
		_, err := tx.Exec("INSERT INTO currencies (code, rate, updated) VALUES ($1, $2, $3) "+
			"ON CONFLICT (code) DO UPDATE SET rate = $2, updated = $3", code, rate, time.Now())
		if err != nil {
			log.Printf("Failed to update currency %s: %v", code, err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
	}
}
