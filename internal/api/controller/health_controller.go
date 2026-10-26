package controller

import (
	"encoding/json"
	"net/http"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "up",
		"message": "Server is running smoothly",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
