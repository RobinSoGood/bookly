package service

import (
	"github.com/RobinSoGood/bookly/internal/domain/models"
	"github.com/RobinSoGood/bookly/internal/logger"
)

type UserStorage interface {
	SaveUser(models.User) (string, error)
	ValidateUser(models.UserLogin) (string, error)
}

type UserService struct {
	stor UserStorage
}

func NewUserService(stor UserStorage) UserService {
	return UserService{stor: stor}
}

func (us *UserService) LoginUser(user models.UserLogin) (string, error) {
	log := logger.Get()
	UID, err := us.stor.ValidateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("validate user failed")
		return ``, err
	}
	return UID, nil
}

func (us *UserService) RegisterUser(user models.User) (string, error) {
	log := logger.Get()
	UID, err := us.stor.SaveUser(user)
	if err != nil {
		log.Error().Err(err).Msg("save user failed")
		return ``, err
	}
	return UID, nil
}
