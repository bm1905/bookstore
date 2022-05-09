package app

import (
	"github.com/bm1905/bookstore_users_api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("starting the application")
	router.Run(":8080")
}
