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

	// Locaations
	mux.HandleFunc("GET /api/locations", getCountriesHandler)
	mux.HandleFunc("GET /api/locations/{code}", getCountryHandler)
	mux.HandleFunc("POST /api/locations", middlewares.Authenticate(createLocationHandler))
	mux.HandleFunc("PUT /api/locations/{code}", middlewares.Authenticate(updateLocationHandler))
	mux.HandleFunc("DELETE /api/locations/{code}", middlewares.Authenticate(deleteLocationHandler))

	// Categories & SubCategories
	mux.HandleFunc("POST /api/categories", middlewares.Authenticate(createCategoryHandler))
	mux.HandleFunc("PUT /api/categories/{id}", middlewares.Authenticate(updateCategoryHandler))
	mux.HandleFunc("DELETE /api/categories/{id}", middlewares.Authenticate(deleteCategoryHandler))
	mux.HandleFunc("POST /api/subcategories", middlewares.Authenticate(createSubCategoryHandler))
	mux.HandleFunc("PUT /api/subcategories/{id}", middlewares.Authenticate(updateSubCategoryHandler))
	mux.HandleFunc("DELETE /api/subcategories/{id}", middlewares.Authenticate(deleteSubCategoryHandler))
	mux.HandleFunc("GET /api/categories-subcategories", middlewares.Authenticate(getCategoriesAndSubcategoriesHandler))

	// Services
	mux.HandleFunc("POST /api/services", middlewares.Authenticate(createServiceHandler))
	mux.HandleFunc("GET /api/services/{id}", getServiceHandler)
	mux.HandleFunc("PUT /api/services/{id}", middlewares.Authenticate(updateServiceHandler))
	mux.HandleFunc("DELETE /api/services/{id}", middlewares.Authenticate(deleteServiceHandler))
	mux.HandleFunc("GET /api/services/{country_code}/{division_id}/{district_id}/{subdistrict_id}/{category_id}/{subcategory_id}", getFilteredServicesHandler)

	return mux
}
