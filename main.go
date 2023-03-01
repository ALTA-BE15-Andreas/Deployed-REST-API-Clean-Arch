package main

import (
	"clean-arch/app/config"
	"clean-arch/app/database"
	"clean-arch/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	database.InitialMigration(db)
	// db := database.InitDBPosgres(*cfg)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	router.InitRouter(db, e)

	e.Logger.Fatal(e.Start(":8080"))
}
