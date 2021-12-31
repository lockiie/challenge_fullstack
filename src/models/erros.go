package models

type Errors struct {
	Message string `json:"message"`
}

func SendError(message string) Errors {
	return Errors{message}
}

type RequestError struct {
	StatusCode int
	Err        error
}

func (e *RequestError) Error() string {
	return e.Err.Error()
}
