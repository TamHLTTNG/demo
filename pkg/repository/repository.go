package repository

import "github.com/Nerzal/gocloak/v13"

type AuthRepository interface {
	GetAccessToken(authorizationCode string) (*gocloak.JWT, error)
	GetIntroSpect(token string) (gocloak.IntroSpectTokenResult, error)
	GetUserInfo(token string) (gocloak.UserInfo, error)
	CreateLocalJWT(information map[string]interface{}) (string, error)
}
