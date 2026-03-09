package calculationService

import "github.com/stretchr/testify/mock"

//поддельный репозиторий
type MockCalcUlationRepository struct {
	mock.Mock
}

func (m *MockCalcUlationRepository) CreateCalculation(calc Calculation) error {
	args := m.Called(calc)
	return args.Error(0)
}

func (m *MockCalcUlationRepository) GetAllCalculation() ([]Calculation, error) {
	args := m.Called()
	var calc []Calculation
	if res := args.Get(0); res != nil {
		calc = res.([]Calculation)

	}
	return calc, args.Error(1)
}

func (m *MockCalcUlationRepository) GetCalculationByID(id string) (Calculation, error) {
	args := m.Called(id)
	var c Calculation
	if res := args.Get(0); res != nil {
		c = res.(Calculation)
	}
	return c, args.Error(1)
}

func (m *MockCalcUlationRepository) UpdateCalculation(calc Calculation) error {
	args := m.Called(calc)
	return args.Error(0)
}

func (m *MockCalcUlationRepository) DeleteCalculation(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
