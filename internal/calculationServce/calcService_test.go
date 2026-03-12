package calculationService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCalculation(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		UserID     string
		mockSetup  func(m *MockCalcUlationRepository, expression string, UserID string)
		wantErr    bool
	}{
		{
			name:       "успешное создание выражения",
			expression: "55+55",
			UserID:     "123",
			mockSetup: func(m *MockCalcUlationRepository, expression string, UserID string) {

				m.On("CreateCalculation", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name:       "ошибка при создании выражения",
			expression: "55+55",
			UserID:     "123",
			mockSetup: func(m *MockCalcUlationRepository, expression string, UserID string) {
				m.On("CreateCalculation", mock.Anything).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, cc := range tests {
		t.Run(cc.name, func(t *testing.T) {
			mockRepo := new(MockCalcUlationRepository)
			cc.mockSetup(mockRepo, cc.expression, cc.UserID)

			service := NewCalculationService(mockRepo)
			result, err := service.CreateCalculation(cc.expression, cc.UserID)

			if cc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, cc.expression, result.Expression)
			}
			mockRepo.AssertExpectations(t)
		})
	}

}

func TestGetAllCalculation(t *testing.T) {
	test := []struct {
		name      string
		mockSetup func(m *MockCalcUlationRepository)
		wantErr   bool
	}{
		{
			name: "Успешно выдали все имеюще выражения",
			mockSetup: func(m *MockCalcUlationRepository) {
				m.On("GetAllCalculation").Return([]Calculation{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Ошибка при выдач всех выражений",
			mockSetup: func(m *MockCalcUlationRepository) {
				m.On("GetAllCalculation").Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}
	for _, cc := range test {
		t.Run(cc.name, func(t *testing.T) {
			mockRepo := new(MockCalcUlationRepository)
			cc.mockSetup(mockRepo)

			service := NewCalculationService(mockRepo)
			result, err := service.GetAllCalculation()

			if cc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, []Calculation{}, result)
			}
			mockRepo.AssertExpectations(t)
		})
	}

}

func TestUpdateCalculation(t *testing.T) {

	tests := []struct {
		name      string
		id        string
		expr      string
		mockSetup func(m *MockCalcUlationRepository, id string, expr string)
		wantErr   bool
	}{
		{
			name: "успешное обновление",
			id:   "123",
			expr: "5+5",
			mockSetup: func(m *MockCalcUlationRepository, id string, expr string) {

				oldCalc := Calculation{
					ID:         id,
					Expression: "2+2",
					Result:     "4",
				}

				m.On("GetCalculationByID", id).Return(oldCalc, nil)

				m.On("UpdateCalculation", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка получения из базы",
			id:   "123",
			expr: "5+5",
			mockSetup: func(m *MockCalcUlationRepository, id string, expr string) {
				m.On("GetCalculationByID", id).
					Return(Calculation{}, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name: "ошибка обновления",
			id:   "123",
			expr: "5+5",
			mockSetup: func(m *MockCalcUlationRepository, id string, expr string) {

				oldCalc := Calculation{
					ID:         id,
					Expression: "2+2",
					Result:     "4",
				}
				m.On("GetCalculationByID", id).Return(oldCalc, nil)
				m.On("UpdateCalculation", mock.Anything).
					Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCalcUlationRepository)
			tt.mockSetup(mockRepo, tt.id, tt.expr)
			service := NewCalculationService(mockRepo)
			result, err := service.UpdateCalculation(tt.id, tt.expr)

			if tt.wantErr {
				assert.Error(t, err)

			} else {

				assert.NoError(t, err)
				assert.Equal(t, tt.expr, result.Expression)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteCalculation(t *testing.T) {
	test := []struct {
		name      string
		id        string
		mockSetup func(m *MockCalcUlationRepository, id string)
		wantErr   bool
	}{
		{
			name: "удаление выполнено",
			id:   "123",
			mockSetup: func(m *MockCalcUlationRepository, id string) {
				m.On("DeleteCalculation", id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "удалить не получилось",
			id:   "123",
			mockSetup: func(m *MockCalcUlationRepository, id string) {
				m.On("DeleteCalculation", id).Return(errors.New("bd error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCalcUlationRepository)
			tt.mockSetup(mockRepo, tt.id)
			service := NewCalculationService(mockRepo)
			err := service.DeleteCalculation(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetCalculationByID(t *testing.T) {
	test := []struct {
		name      string
		id        string
		mockSetup func(m *MockCalcUlationRepository, id string)
		wantErr   bool
	}{
		{
			name: "Успешное нахождение выражения по ID",
			id:   "123",
			mockSetup: func(m *MockCalcUlationRepository, id string) {
				m.On("GetCalculationByID", id).Return(Calculation{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Ошибка при поиске выражения",
			id:   "123",
			mockSetup: func(m *MockCalcUlationRepository, id string) {
				m.On("GetCalculationByID", id).Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCalcUlationRepository)
			tt.mockSetup(mockRepo, tt.id)
			service := NewCalculationService(mockRepo)
			result, err := service.GetCalculationByID(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, Calculation{}, result)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
