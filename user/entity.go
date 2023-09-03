package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID `gorm:"type:char(36);primary_key"`
	Name           string    `gorm:"type:varchar(255)"`
	Occupation     string    `gorm:"type:varchar(255)"`
	Email          string    `gorm:"type:varchar(255);unique"`
	PasswordHash   string    `gorm:"type:varchar(255)"`
	AvatarFileName string    `gorm:"type:varchar(255)"`
	Role           string    `gorm:"type:varchar(255)"`
	Token          string    `gorm:"type:varchar(255)"`
	CreatedAt      time.Time `gorm:"type:datetime"`
	UpdatedAt      time.Time `gorm:"type:datetime"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})

	if err != nil {
		return err
	}

	return nil
}
