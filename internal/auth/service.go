package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"http/test/internal/user"
	"http/test/pkg/di"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	userInstance := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	_, err = service.UserRepository.Create(userInstance)

	if err != nil {
		return "", err
	}

	return userInstance.Email, nil

}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))

	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	return existedUser.Email, nil
}
