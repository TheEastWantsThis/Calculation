package main

import (
	"log"
	calculationService "project/internal/calculationServce"
	"project/internal/db"
	"project/internal/handlers"
	userservice "project/internal/userService"
	"project/internal/web/calculations"
	"project/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Coud noot connect to DB: %v", err)
	}

	e := echo.New()

	usRepo := userservice.NewUserRepository(database)
	usServ := userservice.NewUserService(usRepo)
	usHanler := handlers.NewUserHandler(usServ)

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHanlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.RequestLogger())

	usstrictHandler := users.NewStrictHandler(usHanler, nil)
	users.RegisterHandlers(e, usstrictHandler)

	calcstrictHandler := calculations.NewStrictHandler(calcHanlers, nil) // тут будет ошибка
	calculations.RegisterHandlers(e, calcstrictHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("Coud not start programm err: %v", err)
	}
}
