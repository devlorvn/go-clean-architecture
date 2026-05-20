package postgres

import (
	"context"
	"go-clean-architecture/internal/author"
	database "go-clean-architecture/internal/infras/database/postgres"
	postgresModel "go-clean-architecture/internal/infras/database/postgres/model"

	"gorm.io/gorm"
)

type Author struct {
	db *gorm.DB
}

func NewAuthor(db *gorm.DB) *Author {
	return &Author{db: db}
}

func (a *Author) dbFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := database.TxFromContext(ctx); ok {
		return tx.WithContext(ctx)
	}
	return a.db.WithContext(ctx)
}

func (a *Author) Create(ctx context.Context, author *author.Author) error {
	newAuthor := postgresModel.Author{
		Name: author.Name,
	}
	err := a.dbFromContext(ctx).Create(&newAuthor).Error
	if err != nil {
		return err
	}
	author.ID = newAuthor.ID
	return nil
}

func (a *Author) GetByID(ctx context.Context, id int64) (*author.Author, error) {
	var authorModel postgresModel.Author
	err := a.dbFromContext(ctx).First(&authorModel, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &author.Author{
		ID:     authorModel.ID,
		Name:   authorModel.Name,
		Active: authorModel.Active,
	}, nil
}

func (a *Author) GetAll(ctx context.Context) ([]*author.Author, error) {
	var authorModels []postgresModel.Author
	err := a.dbFromContext(ctx).Find(&authorModels).Error
	if err != nil {
		return nil, err
	}
	authors := make([]*author.Author, len(authorModels))
	for i, m := range authorModels {
		authors[i] = &author.Author{
			ID:     m.ID,
			Name:   m.Name,
			Active: m.Active,
		}
	}
	return authors, nil
}

func (a *Author) Update(ctx context.Context, author *author.Author) error {
	return a.dbFromContext(ctx).Model(&postgresModel.Author{}).Where("id = ?", author.ID).Updates(map[string]interface{}{
		"name":   author.Name,
		"active": author.Active,
	}).Error
}

func (a *Author) Delete(ctx context.Context, id int64) error {
	return a.dbFromContext(ctx).Delete(&postgresModel.Author{}, id).Error
}
