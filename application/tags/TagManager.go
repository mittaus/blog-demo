package tags

import (
	"example.com/mittaus/blog/domain"
	"github.com/gin-gonic/gin"
)

type TagManager struct {
	TagGet     TagGet
	TagsSearch TagsSearch
	// repository domain.ITagRepository
}

func NewTagManager(repository domain.ITagRepository) *TagManager {

	return &TagManager{
		TagGet:     NewTagGet(repository),
		TagsSearch: NewTagSearch(repository),
	}
}

func (tm TagManager) Get(c *gin.Context) {
	TagGet.Get(tm.TagGet, c)
}

func (tm TagManager) Search(c *gin.Context) {

	TagsSearch.Search(tm.TagsSearch, c)
}
