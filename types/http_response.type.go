package types

type ResponseSuccess struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseSuccess(message string, data interface{}) *ResponseSuccess {
	return &ResponseSuccess{
		Success: true,
		Message: message,
		Data:    data,
	}
}

type ResponseError struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

func NewResponseError(message string, errors ...string) *ResponseError {
	return &ResponseError{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}
