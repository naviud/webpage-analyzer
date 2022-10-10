package responses

type ErrorResponse struct {
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

// NewErrorResponse function is responsible to
// create and return a new ErrorResponse object.
func NewErrorResponse(message string, err error) ErrorResponse {
	return ErrorResponse{
		Message:          message,
		DeveloperMessage: err.Error(),
	}
}
