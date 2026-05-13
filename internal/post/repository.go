package post

type Repository interface {
	GetByID(id int64) (*Post, error)
	GetAll() ([]*Post, error)
	Create(post *Post) error
	Update(post *Post) error
	Delete(id int64) error
}
