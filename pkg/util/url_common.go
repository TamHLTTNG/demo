package util

import "go-auth-poc/pkg/config"

type endpointURL struct {
	Token   string
	Inspect string
	Revoke  string
	User    string
	Session string
}

var OpenAM = endpointURL{
	Token:   "/openam/json/access_token",
	Inspect: "/openam/json/inspect",
	Revoke:  "/openam/json/revoke",
	User:    "/openam/json/users/",
	Session: "/openam/json/sessions/",
}

var KeyCloak = endpointURL{
	Token:   config.C.Server.KeyCloakBaseURL + "/auth/realms/" + config.C.Server.KeyCloakRealm + "/protocol/openid-connect/token",
	Inspect: config.C.Server.KeyCloakBaseURL + "/auth/realms/" + config.C.Server.KeyCloakRealm + "/protocol/openid-connect/token/introspect",
	Revoke:  config.C.Server.KeyCloakBaseURL + "/auth/realms/" + config.C.Server.KeyCloakRealm + "/protocol/openid-connect/revoke",
	User:    config.C.Server.KeyCloakBaseURL + "/auth/realms/" + config.C.Server.KeyCloakRealm + "/protocol/openid-connect/userinfo",
	Session: "",
}
