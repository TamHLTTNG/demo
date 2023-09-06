package usecase

import (
	"go-auth-poc/pkg/repository"
)

type AuthUseCase interface {
	GetAccessToken(token string) (*string, error)
	CreateLocalJWT(information map[string]interface{}) (string, error)
}

type authUseCase struct {
	authRepository repository.AuthRepository
	jwtSecret      []byte
}

func NewAuthUseCase(a repository.AuthRepository, jwtSecret string) AuthUseCase {
	return &authUseCase{a, []byte(jwtSecret)}
}

func (ac *authUseCase) GetAccessToken(authorizationCode string) (*string, error) {
	result, err := ac.authRepository.GetAccessToken(authorizationCode)
	if err != nil {
		return nil, err
	}
	return &result.AccessToken, nil
}

func (ac *authUseCase) CreateLocalJWT(information map[string]interface{}) (string, error) {
	result, err := ac.authRepository.CreateLocalJWT(information)
	if err != nil {
		return "", err
	}
	return result, nil
}
