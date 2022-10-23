package helpers

import "github.com/gin-gonic/gin"

func GetTypeContent(c *gin.Context) string {
	return c.Request.Header.Get("Content_Type")
}
