package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
)

func toPublicDetailsRes(v *pb.UserRes) models.PublicUser {
	return models.PublicUser{
		ID:             v.Id,
		Username:       v.Username,
		Email:          v.Email,
		FirstName:      v.FirstName,
		LastName:       v.LastName,
		Phone:          v.Phone,
		Address:        v.Address,
		Description:    v.Description,
		ResumePDF:      v.ResumePdf,
		ResumeDocx:     v.ResumeDocx,
		IsDownloadable: int(v.Isdownloadable),
	}
}

func (h *handler) getPublicDetails(w http.ResponseWriter, r *http.Request) {
	user, err := h.client.GetPublicDetails(h.ctx, &pb.Empty{})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toPublicDetailsRes(user))
}
