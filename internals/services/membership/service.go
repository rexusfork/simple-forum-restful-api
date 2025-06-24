package membership

import (
	"context"
	
	"simple-forum/internals/models/membership"
)

type membershipRepository interface {
	GetUser(contex context.Context, email, username string) (*membership.UserModel, error)
	CreateUser(context context.Context, model membership.UserModel) error
}

type service struct {
	membershipRepo membershipRepository
}

func NewService(membershipRepo membershipRepository) *service {
	return &service{
		membershipRepo: membershipRepo,
	}
}