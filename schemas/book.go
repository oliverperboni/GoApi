package schemas

type Book struct {
	Id     uint `gorm:"primaryKey"`
	Name   string
	Author string
	Genre  string
	Pages  int
}
