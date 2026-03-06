package handlers

import (
	"net/http"
	calculationService "project/internal/calculationServce"

	"github.com/labstack/echo/v4"
)

type CalculationsHandlers struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationsHandlers {
	return &CalculationsHandlers{service: s}
}

func (h *CalculationsHandlers) GetCaculate(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "coud not get calcilations"})
	}

	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationsHandlers) PostCaculate(c echo.Context) error {
	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(req.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Coud not create calculation"})
	}

	return c.JSON(http.StatusCreated, calc)

}

func (h *CalculationsHandlers) PatchCalculation(c echo.Context) error {
	id := c.Param("id")
	var r calculationService.CalculationRequest

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedcalculation, err := h.service.UpdateCalculation(id, r.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Coud not update calculation"})
	}

	return c.JSON(http.StatusOK, updatedcalculation)
}

func (h *CalculationsHandlers) DeleteCalculation(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "coud not delete calcilations"})
	}

	return c.NoContent(http.StatusNoContent)
}
