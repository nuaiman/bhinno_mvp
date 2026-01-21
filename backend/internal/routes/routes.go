package routes

import (
	"backend/internal/middlewares"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Healthchecks
	mux.HandleFunc("/api/health-http", healthCheckHandler)

	// Users
	mux.HandleFunc("POST /api/auth/google", googleAuthHandler)
	mux.HandleFunc("POST /api/auth/email", emailAuthHandler)
	mux.HandleFunc("POST /api/auth/refresh", refreshSessionHandler)
	mux.HandleFunc("GET /api/auth/me", middlewares.Authenticate(getCurrentUserHandler))
	mux.HandleFunc("POST /api/auth/logout", middlewares.Authenticate(logoutHandler))
	mux.HandleFunc("GET /api/users/{id}", middlewares.Authenticate(getUserByIDHandler))

	// Services CRUD
	mux.HandleFunc("POST /api/services", middlewares.Authenticate(createServiceHandler))
	mux.HandleFunc("GET /api/services/{id}", getServiceHandler)
	mux.HandleFunc("PUT /api/services/{id}", middlewares.Authenticate(updateServiceHandler))
	mux.HandleFunc("DELETE /api/services/{id}", middlewares.Authenticate(deleteServiceHandler))
	mux.HandleFunc("GET /api/services/filter/{division_id}/{district_id}/{subdistrict_id}/{category_id}/{subcategory_id}", filterServicesHandler)

	return mux
}
