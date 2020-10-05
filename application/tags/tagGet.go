package tags

import (
	"net/http"
	"strconv"

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
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	tags, err := t.repository.Get(id)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	c.JSON(http.StatusOK, tags)
}
