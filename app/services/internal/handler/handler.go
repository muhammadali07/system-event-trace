package handler

import (
	"github.com/muhammadali07/service-grap-go-api/app/services/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}
