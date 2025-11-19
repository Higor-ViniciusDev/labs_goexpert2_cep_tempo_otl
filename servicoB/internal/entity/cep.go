package entity

import (
	"errors"
	"regexp"
	"strings"
)

var ErrorCepInvalido = errors.New("Cep invalido para uso")

type Cep struct {
	CEP        string
	Logradouro string
	Bairro     string
	Localidade string
	UF         string
	Estado     string
}

func NovoCep(cep string) (*Cep, error) {
	if !ValidaCEP(cep) {
		return nil, ErrorCepInvalido
	}

	return &Cep{
		CEP: cep,
	}, nil
}

func ValidaCEP(cep string) bool {
	cep = strings.TrimSpace(cep)
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}
