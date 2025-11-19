package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNovoCep(t *testing.T) {
	cep, err := NovoCep("15771034")

	assert.Nil(t, err, "Não pode haver erro na criação")
	assert.NotEmpty(t, cep, "Objeto cep não pode ser vazio")
}

func TestCepInvalido(t *testing.T) {
	cep, err := NovoCep("157710")

	assert.NotNil(t, err, "Erro de cep invalido é preciso ser disparado")
	assert.Empty(t, cep, "Objeto retornar vazio")
}
