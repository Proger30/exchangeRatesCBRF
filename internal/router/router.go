package router

import (
	"github.com/Proger30/exchangeRatesCBRF/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/currencies", handlers.GetAllCurrencies(db))
	r.GET("/currency/:code", handlers.GetCurrencyByCode(db))
	r.GET("/update", handlers.ForceUpdateCurrencies(db))

	return r
}
