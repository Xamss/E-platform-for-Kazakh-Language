package server

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"xamss.diploma/api"
)

const authUserID = "auth_user_id"

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorationHeader := ctx.GetHeader("Authorization")
		if authorationHeader == "" {
			err := errors.New("authorization header is not set")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -1,
				Message: err.Error(),
			})
			return
		}

		fields := strings.Fields(authorationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header incorrect format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -2,
				Message: err.Error(),
			})
			return
		}

		userID, err := h.srvs.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -3,
				Message: err.Error(),
			})
			return
		}

		ctx.Set(authUserID, userID)
		ctx.Next()
	}
}

func (h *Handler) CORSMiddleware() gin.HandlerFunc {
	log.Printf("Arrived at CORSMiddleware middleware")
	return func(c *gin.Context) {
		log.Printf("Arrived at CORSMiddleware's return value of handler")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Date")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
