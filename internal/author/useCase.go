package author

import (
	"context"
	"go-clean-architecture/internal/post"
	"go-clean-architecture/internal/shared"
)

type UseCase struct {
	repo     Repository
	postRepo post.Repository
	tx       shared.Transaction
}

func NewUseCase(repo Repository, postRepo post.Repository, tx shared.Transaction) *UseCase {
	return &UseCase{repo: repo, postRepo: postRepo, tx: tx}
}

func (u *UseCase) Create(ctx context.Context, name string) (*Author, error) {
	author := &Author{Name: name, Active: true}

	err := u.repo.Create(ctx, author)

	return author, err
}

func (u *UseCase) GetByID(ctx context.Context, id int64) (*Author, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UseCase) GetAll(ctx context.Context) ([]*Author, error) {
	return u.repo.GetAll(ctx)
}

func (u *UseCase) Update(ctx context.Context, id int64, name string) (*Author, error) {
	author, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	author.Name = name

	err = u.repo.Update(ctx, author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (u *UseCase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *UseCase) Deactivate(ctx context.Context, id int64) error {
	return u.tx.Execute(ctx, func(txCtx context.Context) error {
		author, err := u.repo.GetByID(txCtx, id)
		if err != nil {
			return err
		}
		author.Active = false

		err = u.repo.Update(txCtx, author)
		if err != nil {
			return err
		}

		err = u.postRepo.DeactivateByAuthorID(txCtx, id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *UseCase) Activate(ctx context.Context, id int64) error {
	return u.tx.Execute(ctx, func(txCtx context.Context) error {
		author, err := u.repo.GetByID(txCtx, id)
		if err != nil {
			return err
		}
		author.Active = true

		err = u.repo.Update(txCtx, author)
		if err != nil {
			return err
		}
		err = u.postRepo.ActivateByAuthorID(txCtx, id)
		if err != nil {
			return err
		}
		return nil
	})

}
