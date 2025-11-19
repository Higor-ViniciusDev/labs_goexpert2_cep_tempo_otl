package infra

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscarTemperaturaPorEnderecoValido(t *testing.T) {
	ctx := context.Background()
	serviceCep := NovoViaCepService()
	dadosCep, err := serviceCep.BuscarCepViaService(ctx, "15771034")

	assert.Nil(t, err, "Não pode haver error na busca")
	assert.NotEmpty(t, dadosCep, "Não pode ser cep vazio")

	serviceWeather := NovoWeatherService()
	dadosWeather, err := serviceWeather.BuscarTemperaturaPorEndereco(ctx, dadosCep.Localidade, dadosCep.UF)

	assert.Nil(t, err, "Não pode haver error na busca")
	assert.NotEmpty(t, dadosWeather, "Não pode ser cep vazio")
}
