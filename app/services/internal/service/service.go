package service

import (
	"github.com/muhammadali07/service-grap-go-api/app/services/domain/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
