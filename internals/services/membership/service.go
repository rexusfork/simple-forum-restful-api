package membership

import (
	"context"

	"simple-forum/internals/configs"
	"simple-forum/internals/models/membership"
)

type membershipRepository interface {
	GetUser(contex context.Context, email, username string) (*membership.UserModel, error)
	CreateUser(context context.Context, model membership.UserModel) error
}

type service struct {
	cfg *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg: cfg,
		membershipRepo: membershipRepo,
	}
}