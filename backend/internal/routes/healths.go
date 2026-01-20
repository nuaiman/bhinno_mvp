package routes

import (
	"net/http"

	"backend/internal/utils"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, true, "Service is healthy", map[string]string{"status": "ok"})
}
