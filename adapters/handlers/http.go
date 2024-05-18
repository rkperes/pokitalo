package handlers

import (
	"fmt"
	"net/http"
	"path"
)

const (
	prefix = "/api/v1"
)

type HTTP struct {
	mux *http.ServeMux
}

func NewHTTP() *HTTP {
	h := &HTTP{
		mux: &http.ServeMux{},
	}

	h.register()

	return h
}

func (h *HTTP) register() {
	h.handle(http.MethodGet, "/pokemons/{id}", h.GetPokemon)
}

func (h *HTTP) handle(method string, pattern string, f http.HandlerFunc) {
	h.mux.HandleFunc(fmt.Sprintf("%s %s", method, path.Join(prefix, pattern)), f)
}

// ServeHTTP makes HTTP implement http.Handler interface,
// so it may be served.
func (h *HTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *HTTP) GetPokemon(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
