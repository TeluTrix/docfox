package schema

import (
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string
	Password  string
	IsAdmin   bool `gorm:"default:false"`
}
