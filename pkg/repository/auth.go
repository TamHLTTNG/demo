package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"go-auth-poc/pkg/config"
	"go-auth-poc/pkg/util"
	"io"
	"net/http"
	"net/url"

	"github.com/Nerzal/gocloak/v13"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (ar *authRepository) GetAccessToken(authorizationCode string) (*gocloak.JWT, error) {
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", authorizationCode)
	form.Add("client_id", config.C.Server.KeyCloakClientID)
	form.Add("client_secret", config.C.Server.KeyCloakSecret)

	// Encode the form and create a new request body
	bodyReq := bytes.NewBufferString(form.Encode())

	// Create a new HTTP client and POST request
	response, err := http.Post(util.KeyCloak.Token, "application/x-www-form-urlencoded", bodyReq)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var jwt *gocloak.JWT

	// Decode the response body into a JWT object
	json.Unmarshal(body, &jwt)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}

func (ar *authRepository) GetIntroSpect(token string) (gocloak.IntroSpectTokenResult, error) {
	client := gocloak.NewClient(config.C.Server.KeyCloakBaseURL)
	ctx := context.Background()
	rptResult, err := client.RetrospectToken(ctx, token, config.C.Server.KeyCloakClientID, config.C.Server.KeyCloakSecret, config.C.Server.KeyCloakRealm)
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}

	return *rptResult, nil
}

func (ar *authRepository) GetUserInfo(token string) (gocloak.UserInfo, error) {
	client := gocloak.NewClient(config.C.Server.KeyCloakBaseURL)
	ctx := context.Background()
	userInfo, err := client.GetUserInfo(ctx, token, config.C.Server.KeyCloakRealm)
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}

	return *userInfo, nil
}

func (ar *authRepository) CreateLocalJWT(information map[string]interface{}) (string, error) {
	// Create a JWT token with the data.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(information))

	// Sign the token with a secret key.
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
