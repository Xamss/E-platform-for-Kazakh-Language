package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"xamss.diploma/api"
	"xamss.diploma/internal/model"
)

// createUser registration new user
//
//	@Summary      Create user
//	@Description  Create new user
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param req body api.RegisterRequest true "req body"
//
//	@Success      201
//	@Failure      400  {object}  api.Error
//	@Failure      500  {object}  api.Error
//	@Router       /user/register [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req api.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	u := &model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
	}

	err = h.srvs.CreateUser(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	log.Printf("Arrived at loginUser handler")
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	accessToken, err := h.srvs.Login(ctx, req.Username, req.Password)
	if err != nil {
		log.Printf("login error: %s \n", err.Error())
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    accessToken,
	})
}
