package author

import "context"

type Repository interface {
	GetByID(ctx context.Context, id int64) (*Author, error)
	GetAll(ctx context.Context) ([]*Author, error)
	Create(ctx context.Context, author *Author) error
	Update(ctx context.Context, author *Author) error
	Delete(ctx context.Context, id int64) error
}
