package merchant

import (
	"context"
	"errors"

	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/config"

	"github.com/google/uuid"
)

type merchantRepository interface {
	MerchantInsert(Merchant) (string, error)
	FindMerchantByID(string) (Merchant, error)
}

type Service struct {
	conf *config.Config
	repo merchantRepository
}

func NewService(conf *config.Config, repo merchantRepository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s Service) Register(ctx context.Context, name string) (RegisterResponse, error) {
	m := Merchant{
		Name:     name,
		Username: uuid.New().String(),
		Password: uuid.New().String(),
	}
	id, err := s.repo.MerchantInsert(m)
	if err != nil {
		return RegisterResponse{}, err
	}
	return RegisterResponse{ID: id}, nil
}

func (s Service) Information(ctx context.Context, id string) (Merchant, error) {
	if !s.conf.Merchant.Enable {
		return Merchant{}, errors.New("disabled")
	}

	if id == "" {
		return Merchant{}, errors.New("id cannot be empty")
	}

	m, err := s.repo.FindMerchantByID(id)
	if err != nil {
		return Merchant{}, err
	}
	return m, nil
}
