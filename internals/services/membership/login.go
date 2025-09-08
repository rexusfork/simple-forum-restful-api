package membership

import (
	"context"
	"errors"
	"simple-forum/internals/models/membership"
	"simple-forum/pkg/internals/jwt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, request membership.LoginRequest) (string, error) {
	user, err:= s.membershipRepo.GetUser(ctx, request.Email, "")

	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return "", errors.New("email or password is invalid")
	}

	token, err:= jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		return "", err
	}

	return token, nil
}