package errs

import "fmt"

type HttpError struct {
	CodigoErro int
	Mensagem   string
	Err        error
}

func (e *HttpError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%d - %s: %v", e.CodigoErro, e.Mensagem, e.Err)
	}
	return fmt.Sprintf("%d - %s", e.CodigoErro, e.Mensagem)
}

func (e *HttpError) Unwrap() error {
	return e.Err
}

func New(status int, msg string, err error) *HttpError {
	return &HttpError{
		CodigoErro: status,
		Mensagem:   msg,
		Err:        err,
	}
}
