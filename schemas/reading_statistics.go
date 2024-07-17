package schemas

import "time"

type ReadingStatistic struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uint `gorm:"not null"`
	BookID          uint `gorm:"not null"`
	StartedReading  *time.Time
	FinishedReading *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	User            User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Book            Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}
