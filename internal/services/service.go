package services

import "casbin-go_gin/internal/repository"

type Service struct {
	Auth *AuthSRV
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Auth: NewAuthSRV(repos.Auth)}
}
