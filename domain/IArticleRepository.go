package domain

type IArticleRepository interface {
	Get(id int) (*Article, error)
}
