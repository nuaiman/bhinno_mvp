package routes

import (
	"backend/internal/middlewares"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Healthchecks
	mux.HandleFunc("/api/health-http", healthCheckHandler)

	// Locaations
	mux.HandleFunc("GET /api/countries", countriesHandler)
	mux.HandleFunc("GET /api/locations/{code}", countryDataHandler)

	// Users
	mux.HandleFunc("POST /api/auth/authenticate", authHandler)
	mux.HandleFunc("POST /api/auth/refresh", refreshSessionHandler)
	mux.HandleFunc("GET /api/auth/me", middlewares.Authenticate(getCurrentUserHandler))
	mux.HandleFunc("POST /api/auth/logout", middlewares.Authenticate(logoutHandler))
	mux.HandleFunc("GET /api/users/{id}", middlewares.Authenticate(getUserByIDHandler))

	// Categories & SubCategories
	mux.HandleFunc("POST /api/categories", middlewares.Authenticate(createCategoryHandler))
	mux.HandleFunc("PUT /api/categories/{id}", middlewares.Authenticate(updateCategoryHandler))
	mux.HandleFunc("DELETE /api/categories/{id}", middlewares.Authenticate(deleteCategoryHandler))
	mux.HandleFunc("POST /api/subcategories", middlewares.Authenticate(createSubCategoryHandler))
	mux.HandleFunc("PUT /api/subcategories/{id}", middlewares.Authenticate(updateSubCategoryHandler))
	mux.HandleFunc("DELETE /api/subcategories/{id}", middlewares.Authenticate(deleteSubCategoryHandler))
	mux.HandleFunc("GET /api/categories-subcategories", middlewares.Authenticate(getCategoriesAndSubcategoriesHandler))

	return mux
}
