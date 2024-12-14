package schema

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Item struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
