package infra

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Higor-ViniciusDev/servicoB/internal/entity"
	"github.com/Higor-ViniciusDev/servicoB/internal/errs"
	"github.com/Higor-ViniciusDev/servicoB/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type TempHandler struct {
	TempPorCep *usecase.TemperaturaPorCepUseCase
}

func NovoTempHandler(usecase *usecase.TemperaturaPorCepUseCase) *TempHandler {
	return &TempHandler{
		TempPorCep: usecase,
	}
}

func (t *TempHandler) BuscarTemperaturaPorCep(w http.ResponseWriter, r *http.Request) {
	var input usecase.CepInputDTO

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	tracer := otel.Tracer("servicoB")
	ctx, span := tracer.Start(ctx, "BuscarTemperaturaPorCep Handler")
	defer span.End()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	outPut, err := t.TempPorCep.Execute(ctx, input)

	if err != nil {
		var httpErr *errs.HttpError

		switch {
		case errors.Is(err, entity.ErrorCepInvalido):
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return

		case errors.As(err, &httpErr):
			http.Error(w, httpErr.Error(), httpErr.CodigoErro)
			return

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(outPut)
}
