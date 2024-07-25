package api

import (
	"fmt"
	"log/slog"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"

	"mickamy.com/playground/config"
	"mickamy.com/playground/internal/api/router"
	"mickamy.com/playground/internal/infra/store/database"
)

func Run() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_custom} method=${method} uri=${uri} status=${status}\n",
		CustomTimeFormat: "2006-01-02T15:04:05.999Z07:00",
	}))

	e.Use(middleware.Recover())

	db := db()
	router.User(e, db)

	printRoutes(e)

	apiCfg := config.API()
	port := fmt.Sprintf(":%d", apiCfg.Port)
	e.Logger.Fatal(e.Start(port))
}

func db() *gorm.DB {
	dbCfg := config.DB()
	return database.DB(dbCfg)
}

func printRoutes(e *echo.Echo) {
	routes := e.Routes()
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})
	for _, route := range routes {
		if route.Method == echo.RouteNotFound {
			continue
		}
		slog.Info("Route registered.", "method", route.Method, "path", route.Path)
	}
}
