package schemas

type Book struct {
	Id     uint `gorm:"primaryKey"`
	Name   string
	Author string
	Genre  string
	Pages  int
}
type BookResponse struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Pages  int    `json:"pages"`
}

type BookRepository interface {
	CreateBook(book *Book) error
	GetBookByID(id uint) (*Book, error)
	GetBookByName(name string) (*Book, error)
	UpdateBook(book *Book) error
	DeleteBook(id uint) error
}
