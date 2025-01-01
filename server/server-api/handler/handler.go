package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/token"
	"portfolio/utils"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	ctx        context.Context
	client     pb.ApiServiceClient
	tokenMaker *token.JWTMaker
}

func NewHandler(client pb.ApiServiceClient, secretKey string) *handler {
	return &handler{
		ctx:        context.Background(),
		client:     client,
		tokenMaker: token.NewJWTMaker(secretKey),
	}
}

func (h *handler) getUser(w http.ResponseWriter, r *http.Request) {
	c, ok := r.Context().Value(authKey{}).(*token.UserClaims)
	if !ok {
		helpers.ErrJSONResponse(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.client.GetUser(h.ctx, &pb.UserRes{Username: c.Username})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toUserResPb(user))
}

func (h *handler) createUpdateUser(w http.ResponseWriter, r *http.Request) {
	var body models.Users
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.UserRes{
		Username:       body.Username,
		Email:          body.Email,
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		Phone:          body.Phone,
		Address:        body.Address,
		Description:    body.Description,
		ResumePdf:      body.ResumePDF,
		ResumeDocx:     body.ResumeDocx,
		Isdownloadable: int32(body.IsDownloadable),
	}

	if fullPath[1] == "update" {
		payload.Id = body.ID
	} else if fullPath[1] == "add" {
		hashed, _ := utils.HashPassword(body.Password)
		payload.Password = hashed
	}

	_, err := h.client.CreateUpdateUser(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}

func cleanSplit(path string) []string {
	parts := strings.Split(path, "/")

	var cleanedParts []string
	for _, part := range parts {
		if part != "" {
			cleanedParts = append(cleanedParts, part)
		}
	}
	return cleanedParts
}

func (h *handler) toggleOrDelete(w http.ResponseWriter, r *http.Request) {
	var body models.IDReq
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))

	_, err := h.client.ToggleOrDelete(h.ctx, &pb.IdModel{
		Id:    body.ID,
		Model: fullPath[0],
		Type:  fullPath[1],
	})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", fullPath)
}

func (h *handler) loginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginUserReq
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}

	gu, err := h.client.GetUser(h.ctx, &pb.UserRes{Username: body.Username})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.CheckPassword(gu.Password, body.Password); err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(gu.Id, gu.Email, gu.Username, 15*time.Minute)
	if err != nil {
		helpers.ErrJSONResponse(w, "Failed to create access token", http.StatusInternalServerError)
		return
	}

	refreshToken, refreshClaims, err := h.tokenMaker.CreateToken(gu.Id, gu.Email, gu.Username, 24*time.Hour)
	if err != nil {
		helpers.ErrJSONResponse(w, "Failed to create refresh token", http.StatusInternalServerError)
		return
	}

	_, err = h.client.CreateUpdateSession(h.ctx, &pb.SessionRes{
		Id:           refreshClaims.RegisteredClaims.ID,
		UserId:       gu.Id,
		Email:        gu.Email,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(refreshClaims.ExpiresAt.Time),
	})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(
		w,
		&http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			Expires:  accessClaims.RegisteredClaims.ExpiresAt.Time,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			Secure:   true,
		},
	)

	res := models.RenewAccessTokenReq{
		RefreshToken: refreshToken,
	}

	helpers.JSONResponse(w, "", res)
}

// func (h *handler) logoutUser(w http.ResponseWriter, r *http.Request) {
// 	c := r.Context().Value(authKey{}).(*token.UserClaims)
// 	err := h.client.DeleteSession(h.ctx, c.RegisteredClaims.ID)
// 	if err != nil {
// 		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	helpers.JSONResponse(w, "")
// }

func (h *handler) renewAccessToken(w http.ResponseWriter, r *http.Request) {
	var req models.RenewAccessTokenReq
	if err := helpers.BindValidateJSON(w, r, &req); err != nil {
		return
	}

	refreshClaims, err := h.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		helpers.ErrJSONResponse(w, "invalid refresh token", http.StatusUnauthorized)
		return
	}

	session, err := h.client.GetSession(h.ctx, &pb.IDReq{Id: refreshClaims.RegisteredClaims.ID})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.IsRevoked {
		helpers.ErrJSONResponse(w, "session is revoked", http.StatusBadRequest)
		return
	}
	if session.UserId != refreshClaims.UserID {
		helpers.ErrJSONResponse(w, "invalid session", http.StatusBadRequest)
		return
	}
	if time.Now().After(refreshClaims.RegisteredClaims.ExpiresAt.Time) {
		helpers.ErrJSONResponse(w, "refresh token has expired", http.StatusUnauthorized)
		return
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(
		refreshClaims.UserID,
		refreshClaims.Email,
		refreshClaims.Username,
		15*time.Minute,
	)
	if err != nil {
		helpers.ErrJSONResponse(w, "failed to create access token", http.StatusInternalServerError)
		return
	}

	refreshToken, _, err := h.tokenMaker.CreateToken(
		refreshClaims.UserID,
		refreshClaims.Email,
		refreshClaims.Username,
		24*time.Hour,
	)
	if err != nil {
		helpers.ErrJSONResponse(w, "failed to create refresh token", http.StatusInternalServerError)
		return
	}
	_, err = h.client.CreateUpdateSession(h.ctx, &pb.SessionRes{
		Id:           refreshClaims.RegisteredClaims.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(time.Now().Add(24 * time.Hour)),
		CreatedAt:    timestamppb.New(refreshClaims.RegisteredClaims.IssuedAt.Time),
	})
	if err != nil {
		helpers.ErrJSONResponse(w, "failed to update session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			Expires:  accessClaims.RegisteredClaims.ExpiresAt.Time,
			SameSite: http.SameSiteLaxMode,
			HttpOnly: true,
			Secure:   true,
		},
	)

	helpers.JSONResponse(w, "", models.RenewAccessTokenRes{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
	})
}

// func (h *handler) revokeSession(w http.ResponseWriter, r *http.Request) {
// 	c := r.Context().Value(authKey{}).(*token.UserClaims)
// 	err := h.client.RevokeSession(h.ctx, c.RegisteredClaims.ID)
// 	if err != nil {
// 		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	helpers.JSONResponse(w, "")
// }

func (h *handler) uploadImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		helpers.ErrJSONResponse(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		helpers.ErrJSONResponse(w, "Unable to retrieve file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileExt := strings.ToLower(filepath.Ext(fileHeader.Filename))
	var uploadDir, domain string
	if fileExt == ".jpg" || fileExt == ".jpeg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".svg" || fileExt == ".webp" {
		uploadDir = "/images"
		domain = "http://images.atlascelestia.com"
	} else {
		uploadDir = "/files"
		domain = "http://files.atlascelestia.com"
	}

	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		helpers.ErrJSONResponse(w, "Failed to create upload directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	sanitizedFileName := utils.SanitizeFileName(fileHeader.Filename)
	filePath := filepath.Join(uploadDir, sanitizedFileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		helpers.ErrJSONResponse(w, "Failed to create output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		helpers.ErrJSONResponse(w, "Failed to save the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedURL := fmt.Sprintf("%s/%s", domain, sanitizedFileName)

	_, err = h.client.CreateFile(h.ctx, &pb.FileReq{
		Name: sanitizedFileName,
		File: uploadedURL,
	})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"url": uploadedURL,
	}

	helpers.JSONResponse(w, "File uploaded successfully", res)
}

func (h *handler) getFiles(w http.ResponseWriter, r *http.Request) {
	files, err := h.client.GetFiles(h.ctx, &pb.Empty{})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []models.Files
	for _, file := range files.Files {
		resp = append(resp, models.Files{
			ID:        file.Id,
			Name:      file.Name,
			File:      file.File,
			CreatedAt: file.CreatedAt.AsTime(),
		})
	}

	helpers.JSONResponse(w, "", resp)
}

func (h *handler) deleteFiles(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Url string `json:"url"`
	}
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}

	fileURL := body.Url
	if fileURL == "" {
		helpers.ErrJSONResponse(w, "Missing 'url' in JSON payload", http.StatusBadRequest)
		return
	}

	fileName := filepath.Base(fileURL)
	filePath := filepath.Join("/var/www/uploads/files", utils.SanitizeFileName(fileName))

	err := os.Remove(filePath)
	if err != nil {
		helpers.ErrJSONResponse(w, fmt.Sprintf("Failed to delete file: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = h.client.DeleteFile(h.ctx, &pb.IDReq{Id: fileURL})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "File successfully removed.", http.StatusOK)
}

func (h *handler) isPulicRoute(r *http.Request) *pb.Empty {
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))

	var payload pb.Empty
	if fullPath[0] == "public" {
		payload.OnStatus = true
	}

	return &payload
}
