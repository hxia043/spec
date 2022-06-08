package app

import (
	"iam/internal/controller"
	"iam/internal/pkg/jwt"

	"os"
	"path/filepath"

	"iam/internal/db"

	"github.com/gin-gonic/gin"
)

var defaultTimeout int = 60

type App struct {
	engine *gin.Engine
}

func (app *App) Run() error {
	nonAuthGroup := app.engine.Group("/api/iam/v1")
	authGroup := app.engine.Group("/api/iam/v1")

	authGroup.Use(JWTAuth())

	authGroup.GET("/ping", controller.Ping)
	nonAuthGroup.POST("/registry", controller.Registry)
	nonAuthGroup.POST("/login", controller.Login)

	dbPath := filepath.Join(os.Getenv("HOME"), "sqlite3_iam.db")
	if err := db.InitDatabase(dbPath); err != nil {
		return err
	}

	if err := jwt.InitJwtWrapper(defaultTimeout); err != nil {
		return err
	}

	if err := app.engine.Run("0.0.0.0:9100"); err != nil {
		return err
	}

	return nil
}

func New() *App {
	return &App{
		engine: gin.Default(),
	}
}
