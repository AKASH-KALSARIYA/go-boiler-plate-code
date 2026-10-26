package router

import (
	"go-boilerplate/internal/api/controller"
	"net/http"
)

func SetupRouter(
	healthCtrl *controller.HealthController,
	userCtrl *controller.UserController,
) *http.ServeMux {
	mux := http.NewServeMux()

	// Health Check
	mux.HandleFunc("/health", healthCtrl.Check)

	// User Routes (Full CRUD Example)
	mux.HandleFunc("GET /users", userCtrl.GetAll)
	mux.HandleFunc("POST /users", userCtrl.Create)
	mux.HandleFunc("GET /users/{id}", userCtrl.GetByID)
	mux.HandleFunc("PUT /users/{id}", userCtrl.Update)
	mux.HandleFunc("PATCH /users/{id}", userCtrl.Update)
	mux.HandleFunc("DELETE /users/{id}", userCtrl.Delete)
	mux.HandleFunc("DELETE /users/delete-all", userCtrl.DeleteAll)

	return mux
}
