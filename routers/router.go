package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// API routes
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/health", healthCheck)
	}

	return r
}

// Health check
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
