package util

import "github.com/Higor-ViniciusDev/servicoA/internal/errs"

func CepInvalido(err error) *errs.HttpError {
	return errs.New(422, "invalid zipcode", err)
}

func CepNaoEncontrado(err error) *errs.HttpError {
	return errs.New(404, "can not find zipcode", err)
}
