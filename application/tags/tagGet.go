package tags

import (
	"net/http"

	"example.com/mittaus/blog/domain"
	"github.com/gin-gonic/gin"
)

type ITagGet interface {
	Get(c *gin.Context)
}

type TagGet struct {
	repository domain.ITagRepository
}

func NewTagGet(repository domain.ITagRepository) TagGet {
	return TagGet{
		repository,
	}
}

func (t TagGet) Get(c *gin.Context) {
	id := c.Param("id")

	tags, err := t.repository.Search(id)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	c.JSON(http.StatusOK, tags)
}
