package schemas

import "time"

type UserBook struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	BookID    uint   `gorm:"not null"`
	Status    Status //`gorm:"type:enum('read', 'unread', 'reading');not null"`
	Rating    *int   `gorm:"type:int"`
	Review    string
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Book      Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}
