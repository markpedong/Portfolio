package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func (h *handler) getSession(w http.ResponseWriter, r *http.Request) {
	session, err := h.client.GetSessions(h.ctx, &pb.Empty{})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var sessionsStorer []models.Session
	for _, v := range session.Sessions {
		var session models.Session
		session.ID = v.Id
		session.UserID = v.UserId
		session.RefreshToken = v.RefreshToken
		session.Email = v.Email
		session.IsRevoked = v.IsRevoked
		session.CreatedAt = v.CreatedAt.AsTime()
		session.ExpiresAt = utils.TimestampToTimePtr(v.ExpiresAt)
		sessionsStorer = append(sessionsStorer, session)
	}

	helpers.JSONResponse(w, "", sessionsStorer)
}
