package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Proger30/exchangeRatesCBRF/internal/config"
)

func ConnectDB(cfg config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	migrateDB(db)

	return db, nil
}

func migrateDB(db *sqlx.DB) error {
	migrations := []string{
		"CREATE TABLE IF NOT EXISTS currencies (id SERIAL PRIMARY KEY, code VARCHAR(3) NOT NULL UNIQUE, rate NUMERIC NOT NULL, updated TIMESTAMP NOT NULL)",
	}

	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			return err
		}
	}

	return nil
}
