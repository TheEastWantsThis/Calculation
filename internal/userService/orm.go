package userservice

import (
	calculationService "project/internal/calculationServce"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string                           `gorm:"primaryKey" json:"id"`
	Email        string                           `json:"email"`
	Password     string                           `json:"password"`
	Calculations []calculationService.Calculation `gorm:"foreignKey:user_id"`
	CreatedAt    time.Time                        `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time                       `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt                   `gorm:"index"`
}
