package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func newHandler() *Handler {
	return &Handler{}
}

func (h *Handler) setUpRoutes() {
	fmt.Println("Setting Up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter *http.Request) {
		fmt.Fprint(w, "I am alive")
	})
}
