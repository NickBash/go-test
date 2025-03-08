package auth

import (
	"errors"
	"http/test/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	userInstance := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}

	_, err := service.UserRepository.Create(userInstance)

	if err != nil {
		return "", err
	}

	return userInstance.Email, nil

}

//func (service *AuthService) Login(username, password string) (user.User, error) {}
//
//func (service *AuthService) Logout() error {}
