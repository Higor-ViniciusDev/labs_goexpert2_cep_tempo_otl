package entity

import "context"

type CepService interface {
	BuscarCepViaService(ctx context.Context, cep string) (*Cep, error)
}

type WeatheraService interface {
	BuscarTemperaturaPorEndereco(ctx context.Context, cidade string, estadoSigla string) (*Temperatura, error)
}
