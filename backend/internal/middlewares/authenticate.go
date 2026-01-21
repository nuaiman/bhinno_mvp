package middlewares

import (
	"backend/internal/utils"
	"context"
	"net/http"
	"strings"
)

const (
	CtxUserID string = "userID"
	CtxRole   string = "role"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			utils.JSON(w, http.StatusUnauthorized, false, "Unauthorized", nil)
			return
		}

		accessToken := strings.TrimSpace(authHeader[7:])

		userID, role, err := utils.VerifyJWT(accessToken)
		if err != nil {
			utils.JSON(w, http.StatusUnauthorized, false, "Unauthorized", nil)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, CtxUserID, userID)
		ctx = context.WithValue(ctx, CtxRole, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
