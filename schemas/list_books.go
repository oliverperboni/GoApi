package schemas

import "time"

type ListBook struct {
	ID        uint `gorm:"primaryKey"`
	ListID    uint `gorm:"not null"`
	BookID    uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	List      List `gorm:"foreignKey:ListID;constraint:OnDelete:CASCADE"`
	Book      Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}
