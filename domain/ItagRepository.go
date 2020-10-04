package domain

type ITagRepository interface {
	Search(nombre string) ([]*Tag, error)
}
