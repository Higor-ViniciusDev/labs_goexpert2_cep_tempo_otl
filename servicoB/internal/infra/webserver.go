package infra

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HandlerInfo struct {
	Metodo  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Porta    string
	Handlers map[string][]HandlerInfo // path -> HandlerInfo (method + handler)
	Rotas    chi.Router
}

func NovoWebServer(porta string) *WebServer {
	return &WebServer{
		Porta:    porta,
		Handlers: make(map[string][]HandlerInfo),
		Rotas:    chi.NewRouter(),
	}
}

func (w WebServer) RegistrarRota(caminho string, handlerFunc http.HandlerFunc, metodo string) {
	w.Handlers[caminho] = append(w.Handlers[caminho], HandlerInfo{
		Metodo:  metodo,
		Handler: handlerFunc,
	})
}

func (w WebServer) IniciarWebServer() {
	for rota, handlers := range w.Handlers {
		for _, infoHandle := range handlers {
			w.Rotas.Method(infoHandle.Metodo, rota, infoHandle.Handler)
			log.Printf("Registrando na rota %v com o metodo %v", rota, infoHandle.Metodo)
		}
	}
	w.Rotas.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(w.Porta, w.Rotas)
}
