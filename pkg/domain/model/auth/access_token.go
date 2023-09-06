package auth

type TokenVerifyRequest struct {
	Token string `json:"token"`
}

type AuthorizationCodeRequest struct {
	Code string `json:"code"`
}
