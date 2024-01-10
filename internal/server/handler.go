package server

import (
	"xamss.diploma/internal/usecase"
)

type Handler struct {
	srvs usecase.Service
}

func New(srvs usecase.Service) *Handler {
	return &Handler{
		srvs: srvs,
	}
}
