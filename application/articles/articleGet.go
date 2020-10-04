package articles

import (
	"net/http"
	"strconv"

	"example.com/mittaus/blog/domain"
	"github.com/gin-gonic/gin"
)

type IArticleGet interface {
	Get(c *gin.Context)
}

type ArticleGet struct {
	repository domain.IArticleRepository
}

func NewArticleGet(repository domain.IArticleRepository) ArticleGet {
	return ArticleGet{
		repository,
	}
}

func (t ArticleGet) Get(c *gin.Context) {
	// id := c.Param("id")
	idParametro := c.Param("id")
	id, _ := strconv.Atoi(idParametro)
	tags, err := t.repository.Get(id)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	c.JSON(http.StatusOK, tags)
}
