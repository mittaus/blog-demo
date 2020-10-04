package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	INVALID_PARAMS:                  "Error de parámetro de solicitud",
	ERROR_EXIST_TAG:                 "El nombre de la etiqueta ya existe",
	ERROR_EXIST_TAG_FAIL:            "No se pudo obtener la etiqueta existente",
	ERROR_NOT_EXIST_TAG:             "La etiqueta no existe",
	ERROR_GET_TAGS_FAIL:             "No se pudieron obtener todas las etiquetas",
	ERROR_COUNT_TAG_FAIL:            "Error en la etiqueta de estadísticas",
	ERROR_ADD_TAG_FAIL:              "No se pudo agregar la etiqueta",
	ERROR_EDIT_TAG_FAIL:             "No se pudo modificar la etiqueta",
	ERROR_DELETE_TAG_FAIL:           "No se pudo borrar la etiqueta",
	ERROR_EXPORT_TAG_FAIL:           "No se pudo exportar la etiqueta",
	ERROR_IMPORT_TAG_FAIL:           "No se pudo importar la etiqueta",
	ERROR_NOT_EXIST_ARTICLE:         "El articulo no existe",
	ERROR_ADD_ARTICLE_FAIL:          "No se pudo agregar el artículo",
	ERROR_DELETE_ARTICLE_FAIL:       "No se pudo borrar el artículo",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "No se pudo comprobar si el artículo existe",
	ERROR_EDIT_ARTICLE_FAIL:         "No se pudo modificar el artículo",
	ERROR_COUNT_ARTICLE_FAIL:        "Error en el artículo de estadísticas",
	ERROR_GET_ARTICLES_FAIL:         "No se pudieron obtener varios artículos",
	ERROR_GET_ARTICLE_FAIL:          "No se pudo obtener un solo artículo",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "No se pudo generar el póster del artículo.",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Error de autentificación por token",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Error en token de autenticación",
	ERROR_AUTH_TOKEN:                "Token caducado",
	ERROR_AUTH:                      "Token inválido",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "No se pudo guardar la imagen",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "No se pudo comprobar la imagen",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "Verifique el error de imagen, el formato de imagen o el problema de tamaño",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
