package calculationService

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression, user_id string) (Calculation, error)
	GetAllCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calcService{repo: r}
}

func (s *calcService) Calculatexpression(expression string) (string, error) {
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

func (s *calcService) CreateCalculation(expression, user_id string) (Calculation, error) {
	result, err := s.Calculatexpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc := Calculation{
		ID:         uuid.NewString(),
		UserID:     user_id,
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, err
}

func (s *calcService) GetAllCalculation() ([]Calculation, error) {
	return s.repo.GetAllCalculation()
}

func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *calcService) UpdateCalculation(id string, expression string) (Calculation, error) {

	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	result, err := s.Calculatexpression(expression)

	if err != nil {
		return Calculation{}, err
	}

	calc.Expression = expression
	calc.Result = result

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, err
}

func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}
