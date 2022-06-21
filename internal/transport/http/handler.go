package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetUpRoutes() {
	fmt.Println("Setting Up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter *http.Request) {
		fmt.Fprint(w, "I am alive")
	})
}
