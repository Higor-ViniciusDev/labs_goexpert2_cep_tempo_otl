package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Higor-ViniciusDev/servicoA/internal/entity"
	"github.com/Higor-ViniciusDev/servicoA/internal/errs"
	"github.com/Higor-ViniciusDev/servicoA/internal/usecase"
	"github.com/valyala/fastjson"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type ServicoB struct {
}

func NewServicoB() *ServicoB {
	return &ServicoB{}
}

func (s ServicoB) BuscarInformacaoTempPorCep(ctx context.Context, cep string) (*entity.Temperatura, error) {
	_, span := otel.Tracer("servicoA").Start(ctx, "Chamada ao servicoB - BuscarInformacaoTempPorCep")
	defer span.End()

	bodyRequest := usecase.InputCepDTO{
		Cep: cep,
	}

	jsonBody, err := json.Marshal(bodyRequest)
	if err != nil {
		log.Println(err)
	}

	bodyReader := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequestWithContext(ctx, "POST", "http://servicoB:8181/temperatura", bodyReader)
	// propagar o contexto de tracing via headers
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyErr, _ := io.ReadAll(resp.Body)
		return nil, errs.New(resp.StatusCode, fmt.Sprintf("erro ao buscar temperatura no servicoB: %s", string(bodyErr)), nil)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	data, err := p.ParseBytes(body)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return &entity.Temperatura{
		City:  string(data.GetStringBytes("city")),
		TempC: data.GetFloat64("temp_C"),
		TempF: data.GetFloat64("temp_F"),
		TempK: data.GetFloat64("temp_K"),
	}, nil
}
