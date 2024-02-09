package handlers

import (
	"net/http"

	http_calls "github.com/Proger30/exchangeRatesCBRF/internal/http-calls"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetAllCurrencies(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var currencies []http_calls.Currency
		if err := db.Select(&currencies, "SELECT * FROM currencies"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch currencies"})
			return
		}
		c.JSON(http.StatusOK, currencies)
	}
}

func GetCurrencyByCode(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")

		var currency http_calls.Currency
		if err := db.Get(&currency, "SELECT * FROM currencies WHERE code = $1", code); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Currency not found"})
			return
		}

		c.JSON(http.StatusOK, currency)
	}
}

func ForceUpdateCurrencies(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		http_calls.UpdateCurrencies(db)
		c.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
	}
}
