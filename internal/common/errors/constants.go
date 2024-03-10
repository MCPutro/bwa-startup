package errors

import "net/http"

const (
	statusOK      = "OK"
	statusCreated = "Created"

	//statusNoContent           = "No Content"
	statusNotFound             = "Not Found"
	statusUnauthorized         = "Unauthorized"
	statusInternalServerError  = "Internal Server Error"
	statusBadRequest           = "Bad Request"
	statusUnsupportedMediaType = "Unsupported Media Type"

	//custom error message
	statusUserIdNotFound           = "User ID Not Found"
	statusEmailNotFound            = "Email Not Found"
	statusEmailAlreadyRegister     = "Email Already Register"
	statusCampaignNotFound         = "Campaign Not Found"
	statusEmailAndPasswordNotMatch = "email and password not match"
)

var (
	//NoContext = newError(statusNoContent)

	ErrNotFound             = newError(statusNotFound)
	ErrUnauthorized         = newError(statusUnauthorized)
	ErrInternalServerError  = newError(statusInternalServerError)
	ErrBadRequest           = newError(statusBadRequest)
	ErrUnsupportedMediaType = newError(statusUnsupportedMediaType)

	// List custom error message

	ErrUserIdNotFound           = newError(statusUserIdNotFound)
	ErrEmailNotFound            = newError(statusEmailNotFound)
	ErrCampaignNotFound         = newError(statusCampaignNotFound)
	ErrEmailAlreadyRegister     = newError(statusEmailAlreadyRegister)
	ErrEmailAndPasswordNotMatch = newError(statusEmailAndPasswordNotMatch)
)

var statusMessage = map[string]int{
	statusOK:      200, // StatusOK
	statusCreated: 201, // StatusCreated
	//statusNoContent: 204, // StatusNoContent

	statusBadRequest:           400, // StatusBadRequest
	statusUnauthorized:         401, // StatusUnauthorized
	"Forbidden":                403, // StatusForbidden
	statusNotFound:             404, // StatusNotFound
	"Method Not Allowed":       405, // StatusMethodNotAllowed
	statusUnsupportedMediaType: 415, // Unsupported type

	statusInternalServerError: 500, // StatusInternalServerError

	//custom
	statusUserIdNotFound:           404, // error when user id not found
	statusEmailNotFound:            404, // error when email not found
	statusCampaignNotFound:         404, // error when campaign not found
	statusEmailAlreadyRegister:     400, // error when email already register
	statusEmailAndPasswordNotMatch: 401, // error when wrong password or email

	"test": http.StatusUnauthorized,
}
