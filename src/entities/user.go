package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	BaseModel
	Name        string `gorm:"size:255;not null;" json:"name"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	Password    string `gorm:"size:255;not null;" json:"password"`
	AccountType string `gorm:"size:12;not null;default:'FREE'" json:"account_type"`
	Active      bool   `gorm:"size:1;not null;default:FALSE" json:"active"`
	Language    string `gorm:"size:12;not null;default:'eng';" json:"language"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
