package schemas

import "time"

type Review struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"not null"`
	BookID     uint `gorm:"not null"`
	Rating     int  `gorm:"type:int;not null"`
	ReviewText string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Book       Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}
