package service

import "github.com/MIHAIL33/Service-Nats-streaming/pkg/repository"

type Model interface {

}

type Service struct {
	Model
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}