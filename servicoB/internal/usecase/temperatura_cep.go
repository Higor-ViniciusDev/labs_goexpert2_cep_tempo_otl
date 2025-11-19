package usecase

import (
	"context"

	"github.com/Higor-ViniciusDev/servicoB/internal/entity"
)

type TemperaturaPorCepUseCase struct {
	ServiceCep     entity.CepService
	ServiceWeather entity.WeatheraService
}

func NovoTemperaturaUseCase(serviceCep entity.CepService, serviceWeather entity.WeatheraService) *TemperaturaPorCepUseCase {
	return &TemperaturaPorCepUseCase{
		ServiceCep:     serviceCep,
		ServiceWeather: serviceWeather,
	}
}

type CepInputDTO struct {
	Cep string `json:"cep"`
}

type TemperaturaOutputDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
	City       string  `json:"city,omitempty"`
}

// Execute now accepts a context so it can create internal spans and
// propagate the tracing context to downstream services (viacep/weather).
func (u *TemperaturaPorCepUseCase) Execute(ctx context.Context, input CepInputDTO) (*TemperaturaOutputDTO, error) {
	cepEnt, err := entity.NovoCep(input.Cep)

	if err != nil {
		return nil, err
	}

	// Propagate context to cep service implementation
	cep, err := u.ServiceCep.BuscarCepViaService(ctx, cepEnt.CEP)
	if err != nil {
		return nil, err
	}

	// Propagate context to weather service implementation
	temp, err := u.ServiceWeather.BuscarTemperaturaPorEndereco(ctx, cep.Localidade, cep.UF)

	if err != nil {
		return nil, err
	}
	temp.ConverterCelsiusParaKelvin()

	return &TemperaturaOutputDTO{
		Fahrenheit: temp.TempFar,
		Celsius:    temp.TempCelsius,
		Kelvin:     temp.TempKelvin,
		City:       cep.Localidade,
	}, nil
}
