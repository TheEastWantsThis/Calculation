package handlers

import (
	"context"
	calculationService "project/internal/calculationServce"
	"project/internal/web/calculations"
)

type CalculationsHandlers struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationsHandlers {
	return &CalculationsHandlers{service: s}
}

func (h *CalculationsHandlers) GetCalculations(ctx context.Context, request calculations.GetCalculationsRequestObject) (calculations.GetCalculationsResponseObject, error) {
	calcu, err := h.service.GetAllCalculation()
	if err != nil {
		return nil, err
	}
	response := calculations.GetCalculations200JSONResponse{}

	for _, clc := range calcu {
		calc := calculations.Calculation{
			Id:         &clc.ID,
			Expression: &clc.Expression,
			Result:     &clc.Result,
		}
		response = append(response, calc)
	}

	return response, nil
}

func (h *CalculationsHandlers) PostCalculations(ctx context.Context, request calculations.PostCalculationsRequestObject) (calculations.PostCalculationsResponseObject, error) {
	calc := request.Body

	calcul, err := h.service.CreateCalculation(*calc.Expression)
	if err != nil {
		return nil, err
	}

	response := calculations.PostCalculations201JSONResponse{
		Id:         &calcul.ID,
		Expression: &calcul.Expression,
		Result:     &calcul.Result,
	}

	return response, nil
}

func (h *CalculationsHandlers) PatchCalculationsId(ctx context.Context, request calculations.PatchCalculationsIdRequestObject) (calculations.PatchCalculationsIdResponseObject, error) {
	id := request.Id
	calc := request.Body

	updatedcalculation, err := h.service.UpdateCalculation(id, *calc.Expression)
	if err != nil {
		return nil, err
	}

	response := calculations.PatchCalculationsId200JSONResponse{
		Id:         &updatedcalculation.ID,
		Expression: &updatedcalculation.Expression,
		Result:     &updatedcalculation.Result,
	}

	return response, nil
}

func (h *CalculationsHandlers) DeleteCalculationsId(ctx context.Context, request calculations.DeleteCalculationsIdRequestObject) (calculations.DeleteCalculationsIdResponseObject, error) {

	id := request.Id
	if err := h.service.DeleteCalculation(id); err != nil {
		return nil, err
	}
	return calculations.DeleteCalculationsId204Response{}, nil

}
