package util

const (
	TypeKey                  = "type"
	ErrorKey                 = "error"
	English                  = "en"
	Japan                    = "ja"
	InvalidParameterErrorMsg = "InvalidParameterErrorMsg"
	UnauthorizedErrorMsg     = "UnauthorizedErrorMsg"
	ForbiddenErrorMsg        = "ForbiddenErrorMsg"
	NotSupportedErrorMsg     = "NotSupportedErrorMsg"
	MethodNotAllowedErrorMsg = "MethodNotAllowedErrorMsg"
	DataConflictErrorMsg     = "DataConflictErrorMsg"
	UserDataConflictMsg      = "UserDataConflictMsg"
	InternalServerErrorMsg   = "InternalServerErrorMsg"
	ServiceUnavailableMsg    = "ServiceUnavailableMsg"
	InvalidLanguage          = "InvalidLanguage"
	InvalidKey               = "InvalidKey"
	MissingToken             = "MissingToken"
	TokenInvalid             = "TokenInvalid"
	CallAdmin                = "Cannot login with IBM account. Please contact administrator"
	GetDataFailed            = "Getting data is failed"
	PostRevokeToken          = "Revoke token unsuccessfully"
)

var EnglishList = map[string]string{
	InvalidParameterErrorMsg: "Sending requests with malformed request syntax, invalid request message framing, or deceptive request routing",
	UnauthorizedErrorMsg:     "Unauthorized for access server",
	ForbiddenErrorMsg:        "Unauthorized request",
	NotSupportedErrorMsg:     "Server cannot find the requested resource",
	MethodNotAllowedErrorMsg: "Method Not Allowed.",
	DataConflictErrorMsg:     "Uploading a file that is older than the existing one on the server, resulting in a version control conflict.",
	UserDataConflictMsg:      "User is registered on the server.",
	InternalServerErrorMsg:   "Server encountered an unexpected condition that prevented it from fulfilling the request.",
	ServiceUnavailableMsg:    "Service Unavailable",
	InvalidLanguage:          "Invalid Language",
	InvalidKey:               "Invalid key message",
	MissingToken:             "missing access token",
	TokenInvalid:             "access token is invalid",
}

var JapanList = map[string]string{
	InvalidParameterErrorMsg: "JP - Sending requests with malformed request syntax, invalid request message framing, or deceptive request routing",
	UnauthorizedErrorMsg:     "JP - Unauthorized for access server",
	ForbiddenErrorMsg:        "JP - Unauthorized request",
	NotSupportedErrorMsg:     "JP - Server cannot find the requested resource",
	MethodNotAllowedErrorMsg: "JP - Method Not Allowed.",
	DataConflictErrorMsg:     "JP - Uploading a file that is older than the existing one on the server, resulting in a version control conflict.",
	UserDataConflictMsg:      "JP - User is registered on the server.",
	InternalServerErrorMsg:   "JP - Server encountered an unexpected condition that prevented it from fulfilling the request.",
	ServiceUnavailableMsg:    "JP - Service Unavailable",
	MissingToken:             "JP - missing access token",
	TokenInvalid:             "JP - access token is invalid",
}

func GetList(lang string) (map[string]string, error) {
	switch lang {
	case English:
		return EnglishList, nil
	case Japan:
		return JapanList, nil
	default:
		return nil, &InvalidParameterError{Message: EnglishList[InvalidLanguage]}
	}
}

func GetString(key string) (string, error) {
	list := instance

	value, ok := list[key]
	if !ok {
		return "", &InvalidLanguageError{}
	}

	return value, nil
}

var instance map[string]string

func GetInstance(language string) (map[string]string, error) {
	switch language {
	case English:
		return EnglishList, nil
	case Japan:
		return JapanList, nil
	default:
		return nil, &InvalidParameterError{Message: EnglishList[InvalidKey]}
	}
}
