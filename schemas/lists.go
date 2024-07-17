package schemas

import "time"

type List struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Name      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ListBooks []ListBook `gorm:"foreignKey:ListID"`
	User      User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
