package implementations

import (
	"go-backoffice-seller-api/src/entities"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	GetUserById(id string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	// GetAllUsers() (*[]entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
	// UpdateUser(user *entities.User) (*entities.User, error)
	// DeleteUser(user *entities.User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) GetUserById(id string) (*entities.User, error) {
	var user entities.User
	// result := userRepository.DB.First(&user, id)
	result := userRepository.DB.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (userRepository *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := userRepository.DB.Where("email = ?", email).Find(&user)
	return &user, result.Error
}

// func (userRepository *userRepository) GetAllUsers() (*[]entities.User, error) {
// 	var user []entities.User
// 	result := userRepository.DB.Find(&user)
// 	return &user, result.Error
// }

func (userRepository *userRepository) CreateUser(user *entities.User) (*entities.User, error) {
	result := userRepository.DB.Create(user)
	return user, result.Error
}

// func (userRepository *userRepository) UpdateUser(user *entities.User) (*entities.User, error) {
// 	result := userRepository.DB.Save(user)
// 	return user, result.Error
// }

// func (userRepository *userRepository) DeleteUser(user *entities.User) error {
// 	result := userRepository.DB.Delete(user)
// 	return result.Error
// }
