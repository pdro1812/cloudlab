// currency-converter/internal/api/handler.go
package api

import (
	"encoding/json"
	"net/http"

	"currency-converter/internal/domain"
	"currency-converter/internal/service"
	"github.com/go-chi/chi/v5"
)

type ConversionHandler struct {
	service service.ConversionService
}

// NewConversionHandler cria um novo handler.
// Novamente, note a injeção da INTERFACE do serviço.
func NewConversionHandler(s service.ConversionService) *ConversionHandler {
	return &ConversionHandler{service: s}
}

// RegisterRoutes registra as rotas de conversão no roteador Chi.
func (h *ConversionHandler) RegisterRoutes(r *chi.Mux) {
	r.Post("/convert", h.handleConvert)
}

func (h *ConversionHandler) handleConvert(w http.ResponseWriter, r *http.Request) {
	var req domain.ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validação básica
	if req.From == "" || req.To == "" || req.Amount <= 0 {
		http.Error(w, "Invalid input: from, to and amount are required", http.StatusBadRequest)
		return
	}

	conversion, err := h.service.Convert(r.Context(), req)
	if err != nil {
		// Em um app real, seria bom ter um tratamento de erro mais granular aqui.
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(conversion)
}
