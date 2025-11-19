package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var temp Temperatura

func TestConverterCelsiusParaKelvin(t *testing.T) {
	temp.TempCelsius = 25.5
	temp.TempFar = 45.2

	temp.ConverterCelsiusParaKelvin()

	assert.NotNil(t, temp.TempKelvin, "Após conversão não pode ser vazio temperatura kelvin")
}
