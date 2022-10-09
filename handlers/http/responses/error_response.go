package responses

type ErrorResponse struct {
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

func NewErrorResponse(message string, err error) ErrorResponse {
	return ErrorResponse{
		Message:          message,
		DeveloperMessage: err.Error(),
	}
}
