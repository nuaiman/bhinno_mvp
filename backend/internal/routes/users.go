package routes

import (
	"backend/internal/middlewares"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/auth/credentials/idtoken"
)

func googleAuthHandler(w http.ResponseWriter, r *http.Request) {
	type GoogleAuthRequest struct {
		IDToken     string `json:"id_token"`
		AccessToken string `json:"access_token"`
	}

	var req GoogleAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	if req.IDToken == "" || req.AccessToken == "" {
		utils.JSON(w, http.StatusBadRequest, false, "id_token and access_token required", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	payload, err := idtoken.Validate(ctx, req.IDToken, "YOUR_GOOGLE_CLIENT_ID_HERE")
	if err != nil {
		utils.JSON(w, http.StatusUnauthorized, false, "invalid Google ID token", nil)
		return
	}

	if _, err := utils.VerifyGoogleOAuthAccessToken(req.AccessToken); err != nil {
		utils.JSON(w, http.StatusUnauthorized, false, "invalid Google access token", nil)
		return
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	picture, _ := payload.Claims["picture"].(string)
	googleID, _ := payload.Claims["sub"].(string)

	user, err := models.GetUserByGoogleID(ctx, googleID)
	if err != nil {
		user, err = models.GetUserByEmail(ctx, email)
		if err != nil {
			user = &models.User{
				Email:    email,
				Name:     name,
				Avatar:   picture,
				GoogleID: googleID,
				Verified: true,
			}
			if err := models.CreateUserWithGoogle(ctx, user); err != nil {
				utils.JSON(w, http.StatusInternalServerError, false, "cannot create user", nil)
				return
			}
			user, err = models.GetUserByEmail(ctx, email)
			if err != nil {
				utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch user", nil)
				return
			}
		} else {
			user.GoogleID = googleID
			user.Verified = true
			if user.Avatar == "" {
				user.Avatar = picture
			}
			if err := models.UpdateUser(ctx, user); err != nil {
				utils.JSON(w, http.StatusInternalServerError, false, "cannot link Google account", nil)
				return
			}
		}
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot generate refresh token", nil)
		return
	}
	if err := models.UpdateUserRefreshToken(ctx, user.ID, refreshToken); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot save refresh token", nil)
		return
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot generate access token", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "login successful", map[string]any{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func emailAuthHandler(w http.ResponseWriter, r *http.Request) {
	type AuthRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	if req.Email == "" || req.Password == "" {
		utils.JSON(w, http.StatusBadRequest, false, "email and password required", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := models.GetUserByEmail(ctx, req.Email)
	if err != nil {
		hashedPassword := utils.HashPassword(req.Password)
		user = &models.User{
			Email:    req.Email,
			Password: hashedPassword,
		}
		if err := models.CreateUserWithEmail(ctx, user); err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, "cannot create user", nil)
			return
		}

		user, err = models.GetUserByEmail(ctx, req.Email)
		if err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch user", nil)
			return
		}
	}

	if !utils.CheckHashAndPassword(req.Password, user.Password) {
		utils.JSON(w, http.StatusUnauthorized, false, "invalid credentials", nil)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot generate refresh token", nil)
		return
	}

	if err := models.UpdateUserRefreshToken(ctx, user.ID, refreshToken); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot save refresh token", nil)
		return
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot generate access token", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "login successful", map[string]any{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func refreshSessionHandler(w http.ResponseWriter, r *http.Request) {
	type RefreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := models.GetUserByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		utils.JSON(w, http.StatusUnauthorized, false, "invalid refresh token", nil)
		return
	}

	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot generate refresh token", nil)
		return
	}

	if err := models.UpdateUserRefreshToken(ctx, user.ID, newRefreshToken); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot save refresh token", nil)
		return
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot generate access token", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "token refreshed successfully", map[string]any{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	})
}

func getCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, ok := r.Context().Value(middlewares.CtxUserID).(int64)
	if !ok || userID == 0 {
		utils.JSON(w, http.StatusUnauthorized, false, "unauthorized", nil)
		return
	}

	user, err := models.GetUserByID(ctx, userID)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch user", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "current user fetched", map[string]any{
		"user": user,
	})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userIDVal := r.Context().Value(middlewares.CtxUserID)
	userID, ok := userIDVal.(int64)
	if !ok || userID == 0 {
		utils.JSON(w, http.StatusUnauthorized, false, "unauthorized", nil)
		return
	}

	_ = models.UpdateUserRefreshToken(ctx, userID, "")

	utils.JSON(w, http.StatusOK, true, "logout successful", nil)
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid user ID", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := models.GetUserByID(ctx, userID)
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "user not found", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "user fetched successfully", map[string]any{
		"user": user,
	})
}
