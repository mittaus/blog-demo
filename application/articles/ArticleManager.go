package articles

import (
	"example.com/mittaus/blog/domain"
	"github.com/gin-gonic/gin"
)

type ArticleManager struct {
	ArticleGet ArticleGet
}

func NewArticleManager(repository domain.IArticleRepository) *ArticleManager {

	return &ArticleManager{
		ArticleGet: NewArticleGet(repository),
	}
}

func (am ArticleManager) Get(c *gin.Context) {
	ArticleGet.Get(am.ArticleGet, c)
}
