package calculationService

import (
	"gorm.io/gorm"
)

// основыне методы CRUD -creat, read, update, delete.

type CalculationRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

type calcRepositry struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRepositry{db: db}
}

func (r *calcRepositry) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *calcRepositry) GetAllCalculation() ([]Calculation, error) {
	var calcilations []Calculation
	err := r.db.Find(&calcilations).Error
	return calcilations, err
}

func (r *calcRepositry) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation
	err := r.db.First(calc, "id = ?", id).Error
	return calc, err
}

func (r *calcRepositry) UpdateCalculation(calc Calculation) error {
	err := r.db.Save(&calc).Error
	return err
}

func (r *calcRepositry) DeleteCalculation(id string) error {
	return r.db.Delete(Calculation{}, "id = ?", id).Error
}
