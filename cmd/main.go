package main

import (
	"log"
	calculationService "project/internal/calculationServce"
	"project/internal/db"
	"project/internal/handlers"
	"project/internal/web/calculations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Coud noot connect to DB: %v", err)
	}

	e := echo.New()

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHanlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictHandler := calculations.NewStrictHandler(calcHanlers, nil) // тут будет ошибка
	calculations.RegisterHandlers(e, strictHandler)

	e.Start("localhost:8080")
}
