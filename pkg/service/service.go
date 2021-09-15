package service

import "gitlab.com/alexbizon/info-app/pkg/repository"

type Authorization interface {
}

type InfoList interface {
}

type InfoListItem interface {
}

type Service struct {
	Authorization
	InfoList
	InfoListItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
