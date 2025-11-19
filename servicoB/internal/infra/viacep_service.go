package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Higor-ViniciusDev/servicoB/internal/entity"
	"github.com/Higor-ViniciusDev/servicoB/internal/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type ViaCepService struct{}

func NovoViaCepService() *ViaCepService {
	return &ViaCepService{}
}

func (s *ViaCepService) BuscarCepViaService(ctx context.Context, cep string) (*entity.Cep, error) {
	_, span := otel.Tracer("servicoB").Start(ctx, "BuscarCepViaService - ViaCep API", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Cep        string `json:"cep"`
		Logradouro string `json:"logradouro"`
		Bairro     string `json:"bairro"`
		Localidade string `json:"localidade"`
		Uf         string `json:"uf"`
		Estado     string `json:"estado"`
		Erro       any    `json:"erro"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Erro != nil {
		return nil, util.CepNaoEncontrado(nil)
	}

	return &entity.Cep{
		CEP:        data.Cep,
		Logradouro: data.Logradouro,
		Bairro:     data.Bairro,
		Localidade: data.Localidade,
		UF:         data.Uf,
		Estado:     data.Estado,
	}, nil
}
