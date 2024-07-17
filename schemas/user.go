package schemas

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"size:255;not null"`
	Email        string `gorm:"size:255;not null;unique"`
	PasswordHash string `gorm:"size:255;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserBooks    []UserBook `gorm:"foreignKey:UserID"`
	Lists        []List     `gorm:"foreignKey:UserID"`
	Reviews      []Review   `gorm:"foreignKey:UserID"`
}
