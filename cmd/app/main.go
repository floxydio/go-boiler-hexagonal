package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/floxydio/boiler-fiber/internal/adapters/db"
	"github.com/floxydio/boiler-fiber/internal/adapters/http"
	"github.com/floxydio/boiler-fiber/internal/application"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=root password=rot dbname=test-db port=5432 sslmode=disable"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
		},
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Database Connect")

	repo := db.NewGormProductRepository(gormDB)
	service := application.NewProductService(repo)
	handler := http.NewProductHandler(service)

	app := fiber.New()

	handler.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
