package schemas

type Status string

// gorm:"type:enum('read', 'unread', 'reading');not null"`
const (
	read    Status = "read"
	unread  Status = "unread"
	reading Status = "reading"
)
