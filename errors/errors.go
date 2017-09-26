package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

type AppError struct {
	error
}

type ErrorResult struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}
type ErrorModel struct {
	Code   int         `json:"code,omitempty"`
	Result ErrorResult `json:"error"`
}

func ErrorSQLDuplicateKey(errorString string, uniqueKey string) bool {
	isDuplicateKeyError := strings.Contains(errorString, "duplicate key") && strings.Contains(errorString, uniqueKey)
	if isDuplicateKeyError == false {
		return false
	}
	return true
}

func (e ErrorModel) String() string {
	data, _ := json.MarshalIndent(e, "", "	")
	return string(data)
}
func (e ErrorModel) Error() error {
	// _, file, line, _ := runtime.Caller(1)
	// log.SetPrefix("\x1b[31;1m[Error] ")
	// log.Printf("------------------")
	// log.Printf("Debug %s:%d\n", file, line)
	// log.Printf("%v", m.Meta.ErrorMessage)
	// log.Printf("------------------")
	data, _ := json.MarshalIndent(e, "", "	")
	return errors.New(string(data))
}

func (e ErrorModel) Write(w http.ResponseWriter) {
	codeInt := e.Result.Code
	log.SetPrefix("\x1b[31;1m[Error] ")
	log.Printf("%v %v", e.Result.Code, e.Result.Message)
	http.Error(w, e.String(), codeInt)
}

func (a AppError) Write(w http.ResponseWriter) {
	err := &ErrorModel{}
	unmarshalErr := json.Unmarshal([]byte(a.Error()), &err)
	if unmarshalErr == nil {

		_, file, line, _ := runtime.Caller(1)
		log.SetPrefix("\x1b[31;1m[Error] ")
		log.Printf("------------------")
		log.Printf("Debug %s:%d\n", file, line)
		log.Printf("%v %v\n", err.Result.Code, err.Result.Message)
		log.Printf("------------------")
		codeInt := err.Result.Code
		http.Error(w, fmt.Sprintf("%v", a.Error()), codeInt)

		return
	}
	http.Error(w, fmt.Sprintf("%v", a), http.StatusInternalServerError)
}

type ErrorList interface {
	InvalidUserToken() ErrorModel
	InvalidEmailOrPassword() ErrorModel
	NoPermission() ErrorModel
	NoApplicationKey() ErrorModel
	AdminOnly() ErrorModel
	Unauthorized() ErrorModel
	CannotConnectDatabase() ErrorModel
	UnsupportedCountry() ErrorModel
	UnsupportedMediaType() ErrorModel
	BadRequest(format string, a ...interface{}) ErrorModel
	Teapot(format string, a ...interface{}) ErrorModel
	RecordNotFound(format string, a ...interface{}) ErrorModel
	ParameterRequire(format string, a ...interface{}) ErrorModel
	ServerError(format string, a ...interface{}) ErrorModel
}

func InvalidEmailOrPassword() ErrorModel {
	return NewError(http.StatusBadRequest, "Invalid email or password", http.StatusText(http.StatusBadRequest))
}
func InvalidUserToken() ErrorModel {
	return NewError(http.StatusUnauthorized, "Invalid user token", http.StatusText(http.StatusUnauthorized))
}

func NoPermission() ErrorModel {
	return NewError(http.StatusForbidden, "You do not have permission.", http.StatusText(http.StatusForbidden))
}

func NoApplicationKey() ErrorModel {
	return NewError(http.StatusUnauthorized, "Require Application Id and Rest API key", http.StatusText(http.StatusUnauthorized))
}

func AdminOnly() ErrorModel {
	return NewError(http.StatusForbidden, http.StatusText(http.StatusForbidden), http.StatusText(http.StatusForbidden))
}

func Unauthorized() ErrorModel {
	return NewError(http.StatusUnauthorized, "User required", http.StatusText(http.StatusUnauthorized))
}

func UnsupportedCountry() ErrorModel {
	return NewError(http.StatusTeapot, "Unsupported country", http.StatusText(http.StatusTeapot))
}

func CannotConnectDatabase() ErrorModel {
	return NewError(http.StatusServiceUnavailable, "Cannot connect to database", http.StatusText(http.StatusServiceUnavailable))
}

func UnsupportedMediaType() ErrorModel {
	return NewError(http.StatusBadRequest, "Unsupported media type", http.StatusText(http.StatusBadRequest))
}

func ServerError(format string, a ...interface{}) ErrorModel {
	message := fmt.Sprintf(format, a...)
	return NewError(http.StatusInternalServerError, message, http.StatusText(http.StatusInternalServerError))
}

func BadRequest(format string, a ...interface{}) ErrorModel {
	message := fmt.Sprintf(format, a...)
	return NewError(http.StatusBadRequest, message, http.StatusText(http.StatusBadRequest))
}

func Teapot(format string, a ...interface{}) ErrorModel {
	message := fmt.Sprintf(format, a...)
	return NewError(http.StatusTeapot, message, http.StatusText(http.StatusTeapot))
}

func RecordNotFound(format string, a ...interface{}) ErrorModel {
	message := fmt.Sprintf(format, a...)
	return NewError(http.StatusBadRequest, message, http.StatusText(http.StatusBadRequest))
}

func ParameterRequire(format string, a ...interface{}) ErrorModel {
	message := fmt.Sprintf(format, a...)
	return NewError(http.StatusBadRequest, message, http.StatusText(http.StatusBadRequest))
}
func New(message string) error {
	return NewError(http.StatusInternalServerError, message, http.StatusText(http.StatusInternalServerError)).Error()
}

func NewAppError(message string) AppError {
	return AppError{error: New(message)}
}

func NewAppErrorFromError(e error) AppError {
	return AppError{error: e}
}
func ToAppError(e error) AppError {
	return AppError{error: e}
}

func NewError(code int, errorMessage string, errorType string) ErrorModel {
	return ErrorModel{
		Code: code,
		Result: ErrorResult{
			Code:    code,
			Message: errorMessage,
			Type:    errorType,
		},
	}
}
