package persistence

import (
	"example.com/mittaus/blog/domain"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) domain.IArticleRepository {
	return &ArticleRepository{db}
}

func (repo ArticleRepository) Get(id int) (*domain.Article, error) {
	var article domain.Article
	err := repo.db.First(&article, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}
