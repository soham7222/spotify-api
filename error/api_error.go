package error

import "net/http"

type ErrorResponse struct {
	HttpStatusCode   int       `json:"-"`
	ErrorCode        ErrorCode `json:"error_code"`
	ErrorDescription string    `json:"error_description"`
}

type ErrorCode string

const (
	InternalServerError   ErrorCode = "ERR_SPOTIFY_INTERNAL_SERVER_ERROR"
	BadFormattedJSONError ErrorCode = "ERR_SPOTIFY_BAD_FORMATTED_JSON_ERROR"
	DBInsertionError      ErrorCode = "ERR_SPOTIFY_DB_INSERTION_FAILURE_ERROR"
	DupliacteISRCError    ErrorCode = "ERR_SPOTIFY_DUPLICATE_ISRC_ERROR"
	NoTrackExistsError    ErrorCode = "ERR_SPOTIFY_TRACK_NOT_FOUND"
)

var SpotyfyErrors = map[ErrorCode]*ErrorResponse{
	InternalServerError:   NewErrorResponse(http.StatusInternalServerError, InternalServerError, "internal server error"),
	BadFormattedJSONError: NewErrorResponse(http.StatusBadRequest, BadFormattedJSONError, "malformed json"),
	DBInsertionError:      NewErrorResponse(http.StatusInternalServerError, DBInsertionError, "db insertion failed"),
	DupliacteISRCError:    NewErrorResponse(http.StatusBadRequest, DupliacteISRCError, "duplicate record, isrc already exists"),
	NoTrackExistsError:    NewErrorResponse(http.StatusNotFound, NoTrackExistsError, "no track found"),
}

func NewErrorResponse(statusCode int, errorCode ErrorCode, errorDescription string) *ErrorResponse {
	return &ErrorResponse{
		HttpStatusCode:   statusCode,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
	}
}
