package post

import "context"

type Repository interface {
	GetByID(ctx context.Context, id int64) (*Post, error)
	GetAll(ctx context.Context) ([]*Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int64) error
	DeactivateByAuthorID(ctx context.Context, authorID int64) error
	ActivateByAuthorID(ctx context.Context, authorID int64) error
}
