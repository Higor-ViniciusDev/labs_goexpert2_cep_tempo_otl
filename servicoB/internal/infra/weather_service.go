package infra

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Higor-ViniciusDev/servicoB/internal/entity"
	"github.com/valyala/fastjson"
	"go.opentelemetry.io/otel"
)

type WeatherService struct{}

func NovoWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) BuscarTemperaturaPorEndereco(ctx context.Context, cidade string, estadoSigla string) (*entity.Temperatura, error) {
	_, span := otel.Tracer("servicoB").Start(ctx, "BuscarTemperaturaPorEndereco - WeatherAPI")
	defer span.End()

	cidadeFormatada := url.QueryEscape(cidade)

	urlStr := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=cd5f13c67b234405ab1151712251311&q=%v,%v,brazil&aqi=no", cidadeFormatada, estadoSigla)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
		TempCelsius: data.Get("current").GetFloat64("temp_c"),
		TempFar:     data.Get("current").GetFloat64("temp_f"),
	}, nil
}
