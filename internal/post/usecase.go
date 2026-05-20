package post

import "context"

type UseCase struct {
	Repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{Repo: repo}
}

func (u *UseCase) Create(ctx context.Context, title, content string, authorID int64) (*Post, error) {
	post := &Post{Title: title, Content: content, AuthorID: authorID, Active: true}
	err := u.Repo.Create(ctx, post)

	return post, err
}

func (u *UseCase) GetByID(ctx context.Context, id int64) (*Post, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *UseCase) GetAll(ctx context.Context) ([]*Post, error) {
	return u.Repo.GetAll(ctx)
}

func (u *UseCase) Update(ctx context.Context, id int64, title, content string, active bool) (*Post, error) {
	post, err := u.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	post.Title = title
	post.Content = content
	post.Active = active

	err = u.Repo.Update(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (u *UseCase) Delete(ctx context.Context, id int64) error {
	return u.Repo.Delete(ctx, id)
}
