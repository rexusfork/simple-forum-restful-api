package membership

import (
	"context"
	"errors"
	"time"

	"simple-forum/internals/models/membership"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) Signup(ctx context.Context, req *membership.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Password)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return  err
	}

	now:= time.Now()
	userModel := membership.UserModel{
		Email: req.Email,
		Username: req.Username,
		Password: string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	err = s.membershipRepo.CreateUser(ctx, userModel)
	if err != nil {
		return err
	}

	return nil
}