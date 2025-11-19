package infra

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Higor-ViniciusDev/servicoA/internal/entity"
	"github.com/Higor-ViniciusDev/servicoA/internal/errs"
	"github.com/Higor-ViniciusDev/servicoA/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type getTempHandler struct {
	usecase *usecase.BuscarTempoUsecase
}

func NovoTempHandler(usecase *usecase.BuscarTempoUsecase) *getTempHandler {
	return &getTempHandler{
		usecase: usecase,
	}
}

func (h *getTempHandler) BuscarTemperaturaPorCep(w http.ResponseWriter, r *http.Request) {
	var input usecase.InputCepDTO

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	tracer := otel.Tracer("servicoA")
	ctx, span := tracer.Start(ctx, "BuscarTemperaturaPorCep Handler", trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	outPut, err := h.usecase.Execute(ctx, input)

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
