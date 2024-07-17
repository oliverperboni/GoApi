package schemas

import "time"

type Book struct {
	ID              uint   `gorm:"primaryKey"`
	Title           string `gorm:"size:255;not null"`
	Author          string `gorm:"size:255;not null"`
	Genre           string `gorm:"size:255;not null"`
	PublicationDate *time.Time
	Synopsis        string
	CoverImageURL   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UserBooks       []UserBook `gorm:"foreignKey:BookID"`
	Reviews         []Review   `gorm:"foreignKey:BookID"`
}
