package main

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Calculatexpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), nil

}

func GetCaculate(c echo.Context) error {
	var calculations []Calculation

	if err := db.Find(&calculations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "coud not get calcilations"})
	}
	return c.JSON(http.StatusOK, calculations)
}
func PostCaculate(c echo.Context) error {
	var req CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	res, err := Calculatexpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}
	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: req.Expression,
		Result:     res,
	}
	if err := db.Create(&calc).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "coud not add calcilations"})

	}
	return c.JSON(http.StatusCreated, calc)

}

func PatchCalculation(c echo.Context) error {
	id := c.Param("id")
	var r CalculationRequest
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	res, err := Calculatexpression(r.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	var calc Calculation

	if err := db.First(&calc, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Coud not find exprission"})
	}
	calc.Expression = r.Expression
	calc.Result = res

	if err := db.Save(&calc).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "coud not update calcilations"})
	}
	return c.JSON(http.StatusOK, calc)
}

func DeleteCalculation(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Calculation{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "coud not delete calcilations"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	InitDB()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", GetCaculate)
	e.POST("/calculations", PostCaculate)
	e.PATCH("/calculations/:id", PatchCalculation)
	e.DELETE("/calculations/:id", DeleteCalculation)

	e.Start("localhost:8080")
}
