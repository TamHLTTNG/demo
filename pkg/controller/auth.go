package controller

import (
	"go-auth-poc/pkg/domain/model/auth"
	"go-auth-poc/pkg/repository"
	"go-auth-poc/pkg/usecase"
	"go-auth-poc/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (r *registry) NewAuthController(db *gorm.DB) Auth {
	t := usecase.NewAuthUseCase(
		repository.NewAuthRepository(db),
		"jwt_secret",
	)

	return NewAuthController(t)
}

type Auth interface {
	POSTTokenVerify(c *gin.Context)
}

type authController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(us usecase.AuthUseCase) Auth {
	return &authController{us}
}

// POS TTokenVerify godoc
// @Summary Verify Access Token
// @Description Verify access token is valid
// @Tags authentication
// @Accept  json
// @Produce  json
// @Param information body auth.AuthorizationCodeRequest true "Authorization Code"
// @Success 200 {string} string "Param is valid"
// @Failure 400 {string} string "missing Param"
// @Failure 401 {string} string "Param is invalid"
// @Failure 500 {string} string "Internal Server Error"
// @Router /get-jwt [post]
func (ac *authController) POSTTokenVerify(c *gin.Context) {

	var codeRequest auth.AuthorizationCodeRequest

	if err := c.BindJSON(&codeRequest); err != nil {
		util.HandleError(c, err)
		return
	}

	token, err := ac.authUseCase.GetAccessToken(codeRequest.Code)

	if err != nil {
		util.HandleError(c, err)
		return
	}

	data := map[string]interface{}{
		"token-keycloak": token,
	}

	jwt, err := ac.authUseCase.CreateLocalJWT(data)

	if err != nil {
		c.AbortWithStatusJSON(500, err)
	} else {
		c.JSON(http.StatusOK, map[string]string{
			"token_jwt": jwt,
		})
	}

}
