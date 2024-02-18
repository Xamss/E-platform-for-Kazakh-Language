package server

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(h.CORSMiddleware())

	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")

	user.POST("/register", h.createUser)
	user.POST("/login", h.loginUser)

	//post.Use(h.authMiddleware())
	return router
}
