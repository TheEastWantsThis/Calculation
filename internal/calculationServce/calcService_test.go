package calculationService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCalculation(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		mockSetup func(m *MockCalcUlationRepository, input string)
		wantErr   bool
	}{
		{
			name:  "успешное создание выражения",
			input: "55+55",
			mockSetup: func(m *MockCalcUlationRepository, input string) {

				m.On("CreateCalculation", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при создании выражения",
			input: "55+55",
			mockSetup: func(m *MockCalcUlationRepository, input string) {
				m.On("CreateCalculation", mock.Anything).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, cc := range tests {
		t.Run(cc.name, func(t *testing.T) {
			mockRepo := new(MockCalcUlationRepository)
			cc.mockSetup(mockRepo, cc.input)

			service := NewCalculationService(mockRepo)
			result, err := service.CreateCalculation(cc.input)

			if cc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, cc.input, result.Expression)
			}
			mockRepo.AssertExpectations(t)
		})
	}

}
