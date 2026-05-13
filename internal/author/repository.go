package author

type Repository interface {
	GetByID(id int) (*Author, error)
	GetAll() ([]*Author, error)
	Create(author *Author) error
	Update(author *Author) error
	Delete(id int) error
}
