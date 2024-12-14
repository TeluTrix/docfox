package main

import (
	"github.com/TeluTrix/docfox/api"
	"github.com/TeluTrix/docfox/internal"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func init() {
	LoadEnv()
	internal.ConnectToDB()
	internal.MigrateModels()
}

func main() {
	router := gin.Default()

	api.AuthRoutes(router)
	api.DocRoutes(router)

	err := router.Run()
	if err != nil {
		slog.Error("Could not start server")
		os.Exit(1)
	}
}
