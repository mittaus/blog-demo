package tags

type ITagManager interface {
	ITagGet
	ITagsSearch
	// Search(c *gin.Context)
	// Get(c *gin.Context)
	// Put(c *gin.Context)
}
