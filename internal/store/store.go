package store

const (
	CUSTOMER = iota
	STAFF
	ADMIN
)

type User struct {
	Name     string
	LastName string
	Email    string
	Password string
	Role     int
}

type Store interface {
	Create(*User) error
	Update(*User, int) error
	Delete(int) error
	GetById(int) (*User, error)
}
