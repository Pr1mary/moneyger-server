package main

import (
	"moneyger-server/internal/auth"
	"moneyger-server/internal/health"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	health.Router(r)
	auth.Router(r)
	r.Run(":9000")
}
