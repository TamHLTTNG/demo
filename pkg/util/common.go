package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetInstanceLanguage(c *gin.Context) {
	language := c.GetHeader("Language")
	if language == "" {
		// HandleError(c, &InvalidParameterError{Message: "missing language"})
		language = "en"
	}
	value, err := GetInstance(language)
	if err != nil {
		// HandleError(c, &InvalidParameterError{Message: "language not support"})
		valueEn, _ := GetInstance("en")
		instance = valueEn
		return
	}
	instance = value
}

func HandleError(c *gin.Context, err error) {
	errMsg := MapErrorToAppError(err)
	code := errMsg.error_code
	typeErr := errMsg.error_type
	message := errMsg.message

	c.JSON(code, gin.H{
		TypeKey:  typeErr,
		ErrorKey: message,
	})
}

func ReturnError(err AppError) string {
	return fmt.Sprintf("status_code=%d, error_type=%s, error_code=%d, message=%s", err.status_code, err.error_type, err.error_code, err.message)
}

type AppError struct {
	status_code int
	error_type  string
	error_code  int
	message     string
}

func MapErrorToAppError(err error) AppError {
	if err == nil {
		return AppError{}
	}

	parts := strings.Split(err.Error(), ", ")
	if len(parts) <= 1 {
		return AppError{
			status_code: 500,
			error_type:  "InternalServerError",
			error_code:  500,
			message:     parts[0]}
	}
	status_code, _ := strconv.Atoi(strings.Split(parts[0], "=")[1])
	error_type := strings.Split(parts[1], "=")[1]
	error_code, _ := strconv.Atoi(strings.Split(parts[2], "=")[1])
	message := strings.Split(parts[3], "=")[1]

	return AppError{
		status_code: status_code,
		error_type:  error_type,
		error_code:  error_code,
		message:     message,
	}
}

// Query request error (200)
type QueryError struct {
	Message string
}

func (e *QueryError) Error() string {
	/// 200 ERROR: Query request error

	err := AppError{
		status_code: 200,
		error_type:  "QueryError",
		error_code:  200,
		message:     e.Message}
	return ReturnError(err)
}

// Bad request (400)
type InvalidLanguageError struct {
	Message string
}

func (e *InvalidLanguageError) Error() string {
	/// 400 ERROR: Invalid Parameter (Bad request)

	err := AppError{
		status_code: 400,
		error_type:  "InvalidParameterError",
		error_code:  400,
		message:     e.Message + EnglishList[InvalidLanguage]}
	return ReturnError(err)
}

// Bad request (400)
type InvalidParameterError struct {
	Message string
}

func (e *InvalidParameterError) Error() string {
	/// 400 ERROR: Invalid Parameter (Bad request)

	err := AppError{
		status_code: 400,
		error_type:  "InvalidParameterError",
		error_code:  400,
		message:     e.Message}
	return ReturnError(err)
}

// Bad request (401)
type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	/// 401 ERROR: UnauthorizedError

	errMsg, errLanguage := GetString(UnauthorizedErrorMsg)
	if errLanguage != nil {
		return ReturnError(AppError{status_code: 500, error_code: 500, error_type: "Get Language", message: errLanguage.Error()})
	}

	err := AppError{
		status_code: 401,
		error_type:  "UnauthorizedError",
		error_code:  401,
		message:     errMsg + e.Message}
	return ReturnError(err)
}

// Forbidden (403)
type ForbiddenError struct{}

func (e *ForbiddenError) Error() string {
	/// 403 ERROR: ForbiddenError

	errMsg, errLanguage := GetString(ForbiddenErrorMsg)
	if errLanguage != nil {
		return ReturnError(AppError{status_code: 500, error_code: 500, error_type: "Get Language", message: errLanguage.Error()})
	}
	err := AppError{
		status_code: 403,
		error_type:  "ForbiddenError",
		error_code:  403,
		message:     errMsg}
	return ReturnError(err)
}

// Not found (404)
type NotSupportedError struct{}

func (e *NotSupportedError) Error() string {
	/// 404 ERROR: Unsupported Request

	errMsg, errLanguage := GetString(NotSupportedErrorMsg)
	if errLanguage != nil {
		return ReturnError(AppError{status_code: 500, error_code: 500, error_type: "Get Language", message: errLanguage.Error()})
	}
	err := AppError{
		status_code: 404,
		error_type:  "NotSupportedError",
		error_code:  404,
		message:     errMsg}
	return ReturnError(err)
}

// Method Not Allowed (405)
type MethodNotAllowedError struct {
	Message string
}

func (e *MethodNotAllowedError) Error() string {
	/// 405 ERROR: Method Not Allowed

	errMsg, errLanguage := GetString(MethodNotAllowedErrorMsg)
	if errLanguage != nil {
		return ReturnError(AppError{status_code: 500, error_code: 500, error_type: "Get Language", message: errLanguage.Error()})
	}
	err := AppError{
		status_code: 405,
		error_type:  "MethodNotAllowedError",
		error_code:  405,
		message:     errMsg + e.Message}
	return ReturnError(err)
}

// Data conflict (409)
type DataConflictError struct{}

func (e DataConflictError) Error() string {
	/// 409 ERROR: Data is conflicted

	errMsg, errLanguage := GetString(DataConflictErrorMsg)
	if errLanguage != nil {
		return ReturnError(AppError{status_code: 500, error_code: 500, error_type: "Get Language", message: errLanguage.Error()})
	}
	err := AppError{
		status_code: 409,
		error_type:  "DataConflictError",
		error_code:  409,
		message:     errMsg}
	return ReturnError(err)
}

// Data conflict (409)
type UserDataConflict struct{}

func (e *DataConflictError) UserDataConflict() string {
	/// 409 ERROR: User data is conflicted

	errMsg, errLanguage := GetString(UserDataConflictMsg)
	if errLanguage != nil {
		return ReturnError(AppError{status_code: 500, error_code: 500, error_type: "Get Language", message: errLanguage.Error()})
	}
	err := AppError{
		status_code: 409,
		error_type:  "UserDataConflict",
		error_code:  409,
		message:     errMsg}
	return ReturnError(err)
}

// Service unavailable (500)
type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	/// 503 ERROR: Internal server error

	err := AppError{
		status_code: 500,
		error_type:  "InternalServerError",
		error_code:  500,
		message:     e.Message}
	return ReturnError(err)
}

// Service unavailable (503)
type ServiceUnavailable struct{}

func (e *ServiceUnavailable) Error() string {
	/// 503 ERROR: Service is temporarily unavailable

	errMsg, errLanguage := GetString(ServiceUnavailableMsg)
	if errLanguage != nil {
		return ReturnError(AppError{status_code: 500, error_code: 500, error_type: "Get Language", message: errLanguage.Error()})
	}
	err := AppError{
		status_code: 503,
		error_type:  "ServiceUnavailable",
		error_code:  503,
		message:     errMsg}
	return ReturnError(err)
}
