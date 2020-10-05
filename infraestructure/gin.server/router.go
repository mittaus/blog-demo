package server

import (
	articles "example.com/mittaus/blog/application/articles"
	tags "example.com/mittaus/blog/application/tags"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	tagHandler     tags.ITagManager
	articleHandler articles.IArticleManager
	// ucHandler   application.tags
	// authHandler application.AuthHandler
	// Logger      application.Logger
}

func NewRouterHandler(tagHandler tags.ITagManager, articleHandler articles.IArticleManager) RouterHandler {
	return RouterHandler{
		tagHandler:     tagHandler,
		articleHandler: articleHandler,
	}
}

//SetRoutes configuración inicial
func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	// r.POST("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(authHandler.JWT())
	// {
	// //Obtenga una lista de etiquetas
	apiv1.GET("/tags", rH.tagHandler.Search)
	// //Nueva etiqueta
	// apiv1.POST("/tags", v1.AddTag)
	// //Actualizar etiqueta especificada
	apiv1.GET("/tag/:id", rH.tagHandler.Get)
	// //Eliminar etiqueta especificada
	// apiv1.DELETE("/tags/:id", v1.DeleteTag)
	// //Exportar etiqueta
	// r.POST("/tags/export", v1.ExportTag)
	// //Importar etiquetas
	// r.POST("/tags/import", v1.ImportTag)

	// //Obtenga una lista de artículos
	// apiv1.GET("/articles", v1.GetArticles)
	// //Obtén el artículo especificado
	apiv1.GET("/articles/:id", rH.articleHandler.Get)
	// //Articulo nuevo
	// apiv1.POST("/articles", v1.AddArticle)
	// //Actualizar el artículo especificado
	// apiv1.PUT("/articles/:id", v1.EditArticle)
	// //Eliminar artículo especificado
	// apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	// //Generar póster de artículo
	// apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

	// }

	// return r
}
