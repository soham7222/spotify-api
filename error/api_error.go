package error

import "net/http"

type ErrorResponse struct {
	HttpStatusCode   int         `json:"-"`
	ErrorCode        ErrorCode   `json:"error_code"`
	ErrorDescription string      `json:"error_description"`
	AdditionalData   interface{} `json:"additional_data"`
}

type ErrorCode string

const (
	InternalServerError   ErrorCode = "ERR_SPOTIFY_INTERNAL_SERVER_ERROR"
	BadFormattedJSONError ErrorCode = "ERR_SPOTIFY_BAD_FORMATTED_JSON_ERROR"
)

var SpotyfyErrors = map[ErrorCode]*ErrorResponse{
	InternalServerError:   NewErrorResponse(http.StatusInternalServerError, InternalServerError, "internal server error"),
	BadFormattedJSONError: NewErrorResponse(http.StatusBadRequest, BadFormattedJSONError, "malformed json"),
}

func NewErrorResponse(statusCode int, errorCode ErrorCode, errorDescription string) *ErrorResponse {
	return &ErrorResponse{
		HttpStatusCode:   statusCode,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
	}
}
