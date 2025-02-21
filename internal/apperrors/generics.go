package apperrors

type basicError struct {
	Success        bool   `json:"success"`
	Err            string `json:"error"`
	Cause          string `json:"cause"`
	httpStatusCode int
}

func (e *basicError) Error() string {
	return e.Err + ": " + e.Cause
}

func (e *basicError) SetCause(cause string) *basicError {
	e.Cause = cause

	return e
}

func newError(err string, httpStatusCode int) *basicError {
	return &basicError{
		Err:            err,
		httpStatusCode: httpStatusCode,
	}
}
