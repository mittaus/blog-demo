package tags

import (
	"net/http"

	"example.com/mittaus/blog/domain"
	"github.com/gin-gonic/gin"
)

type ITagsSearch interface {
	Search(c *gin.Context)
}

type TagsSearch struct {
	repository domain.ITagRepository
}

func NewTagSearch(repository domain.ITagRepository) TagsSearch {
	return TagsSearch{
		repository,
	}
}

func (t TagsSearch) Search(c *gin.Context) {
	name, _ := c.GetQuery("name")

	tags, err := t.repository.Search(name)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	c.JSON(http.StatusOK, tags)
}
