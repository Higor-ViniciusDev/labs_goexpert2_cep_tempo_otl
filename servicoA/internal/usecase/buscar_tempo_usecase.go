package usecase

import (
	"context"

	"github.com/Higor-ViniciusDev/servicoA/internal/entity"
)

type BuscarTempoUsecase struct {
	tempoServico entity.ServicoBInterface
}

func NewBuscarTempoUsecase(tempoServico entity.ServicoBInterface) *BuscarTempoUsecase {
	return &BuscarTempoUsecase{
		tempoServico: tempoServico,
	}
}

type InputCepDTO struct {
	Cep string `json:"cep"`
}

// { "city: "São Paulo", "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
type OutputTempoPorCep struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// Ação para receber um cep input validar se ele é ok, e fazer um requisição atraves do serviço da b para retornar os dados
func (b *BuscarTempoUsecase) Execute(ctx context.Context, input InputCepDTO) (*OutputTempoPorCep, error) {
	novoCep, err := entity.NovoCep(input.Cep)
	if err != nil {
		return nil, err
	}

	dadosTemp, err := b.tempoServico.BuscarInformacaoTempPorCep(ctx, novoCep.Cep)

	if err != nil {
		return nil, err
	}

	novoRetorno := &OutputTempoPorCep{
		City:  dadosTemp.City,
		TempC: dadosTemp.TempC,
		TempF: dadosTemp.TempF,
		TempK: dadosTemp.TempK,
	}

	return novoRetorno, nil
}
