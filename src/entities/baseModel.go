package entities

import (
	"time"

	_ "github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
