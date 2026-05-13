package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(201, gin.H{
		"success": true,
		"data":    data,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"success": false,
		"error":   message,
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(401, gin.H{
		"success": false,
		"error":   message,
	})
}

func Forbidden(c *gin.Context, message string) {
	c.JSON(403, gin.H{
		"success": false,
		"error":   message,
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(404, gin.H{
		"success": false,
		"error":   message,
	})
}

func InternalError(c *gin.Context, message string) {
	c.JSON(500, gin.H{
		"success": false,
		"error":   message,
	})
}
