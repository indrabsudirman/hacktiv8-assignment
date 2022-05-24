package repository

import (
	"hacktiv8-day07/sesi07-gorm/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(*models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserByID(uint) (*models.User, error)
	UpdateUserByID(*models.User, uint) (*models.User, error)
	DeleteUserByID(uint) error
	GetUsersWithProducts() (*[]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// DeleteUserByID implements UserRepository
func (ur *userRepository) DeleteUserByID(id uint) error {
	user := models.User{}
	err := ur.db.Delete(&user, "id=?", id).Error
	return err
}

// UpdateUserByID implements UserRepository
func (ur *userRepository) UpdateUserByID(request *models.User, id uint) (*models.User, error) {
	user := models.User{}
	e := ur.db.First(&user, "id=?", id).Error
	if e != nil {
		return nil, e
	}
	err := ur.db.Where("id=?", id).Updates(&request)
	user = *request
	user.ID = id
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

// CreateUser implements UserRepository
func (ur *userRepository) CreateUser(request *models.User) error {
	err := ur.db.Create(request).Error
	return err
}

// GetAllUsers implements UserRepository
func (ur *userRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User

	err := ur.db.Find(&users).Error
	return &users, err
}

// GetUserByID implements UserRepository
func (ur *userRepository) GetUserByID(id uint) (*models.User, error) {
	user := models.User{}

	err := ur.db.First(&user, "id=?", id).Error
	return &user, err
}

func (ur *userRepository) GetUsersWithProducts() (*[]models.User, error) {
	users := []models.User{}
	err := ur.db.Preload("Products").Find(&users).Error
	return &users, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
