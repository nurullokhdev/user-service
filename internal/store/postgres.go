package store

type PostgreStore struct{}

func NewPostgre() *PostgreStore {
	return &PostgreStore{}
}

func (s *PostgreStore) Create(u *User) error {
	//...
	return nil
}

func (s *PostgreStore) Update(u *User, id int) error {
	//...
	return nil
}

func (s *PostgreStore) Delete(id int) error {
	//...
	return nil
}

func (s *PostgreStore) GetById(id int) (*User, error) {
	//...
	return &User{
		Name:     "Alice",
		LastName: "Doe",
		Email:    "alicedoe@suncloud.com",
	}, nil
}
