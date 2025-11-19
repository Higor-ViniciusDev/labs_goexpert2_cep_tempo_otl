package main

import (
	"context"
	"fmt"

	"github.com/Higor-ViniciusDev/servicoB/internal/infra"
	"github.com/Higor-ViniciusDev/servicoB/internal/usecase"
)

func main() {
	ctx := context.Background()
	shutdown, err := infra.InitProviders(ctx, "servicoB")
	if err != nil {
		fmt.Printf("Erro ao inicializar providers de telemetria: %v\n", err)
	} else {
		defer func() {
			_ = shutdown(context.Background())
		}()
	}

	useCaseTemp := usecase.NovoTemperaturaUseCase(infra.NovoViaCepService(), infra.NovoWeatherService())
	novoHandler := infra.NovoTempHandler(useCaseTemp)
	web := infra.NovoWebServer(":8181")
	web.RegistrarRota("/temperatura", novoHandler.BuscarTemperaturaPorCep, "POST")

	web.IniciarWebServer()
}
