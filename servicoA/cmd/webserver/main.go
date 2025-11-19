package main

import (
	"context"
	"fmt"

	"github.com/Higor-ViniciusDev/servicoA/internal/infra"
	"github.com/Higor-ViniciusDev/servicoA/internal/infra/web"
	"github.com/Higor-ViniciusDev/servicoA/internal/service"
	"github.com/Higor-ViniciusDev/servicoA/internal/usecase"
)

func main() {
	ctx := context.Background()
	shutdown, err := infra.InitProviders(ctx, "servicoA")
	if err != nil {
		fmt.Printf("Erro ao inicializar providers de telemetria: %v\n", err)
	} else {
		defer func() {
			_ = shutdown(context.Background())
		}()
	}

	useCaseTemp := usecase.NewBuscarTempoUsecase(service.NewServicoB())
	novoHandler := infra.NovoTempHandler(useCaseTemp)
	web := web.NovoWebServer(":8080")
	web.RegistrarRota("/temperatura", novoHandler.BuscarTemperaturaPorCep, "POST")

	web.IniciarWebServer()
}
