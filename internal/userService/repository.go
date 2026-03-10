package userservice

import "gorm.io/gorm"

type UserRepository interface {
	GetAllUsers() ([]User, error)
	CreateUser(user User) error
	GetUserByID(id string) (User, error)
	UpdateUser(user User) error
	DeleteUser(id string) error
}

type uRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &uRepo{db: db}
}
func (r *uRepo) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *uRepo) GetAllUsers() ([]User, error) {
	var us []User
	err := r.db.Find(&us).Error
	return us, err
}

func (r *uRepo) GetUserByID(id string) (User, error) {
	var us User
	err := r.db.First(&us, "id=?", id).Error
	return us, err
}

func (r *uRepo) UpdateUser(user User) error {
	err := r.db.Save(user).Error
	return err
}

func (r *uRepo) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
