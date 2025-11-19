package entity

import "context"

type ServicoBInterface interface {
	BuscarInformacaoTempPorCep(ctx context.Context, cep string) (*Temperatura, error)
}
