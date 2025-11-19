package entity

import (
	"errors"
	"regexp"
	"strings"
)

var ErrorCepInvalido = errors.New("invalid zipcode")

type Cep struct {
	Cep string
}

func NovoCep(cep string) (*Cep, error) {
	if !validaCEP(cep) {
		return nil, ErrorCepInvalido
	}

	return &Cep{
		Cep: cep,
	}, nil
}

func validaCEP(cep string) bool {
	cep = strings.TrimSpace(cep)
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}
