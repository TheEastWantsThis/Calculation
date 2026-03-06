package main

import (
	"log"
	calculationService "project/internal/calculationServce"
	"project/internal/db"
	"project/internal/handlers"

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

	e.GET("/calculations", calcHanlers.GetCaculate)
	e.POST("/calculations", calcHanlers.PostCaculate)
	e.PATCH("/calculations/:id", calcHanlers.PatchCalculation)
	e.DELETE("/calculations/:id", calcHanlers.DeleteCalculation)

	e.Start("localhost:8080")
}
