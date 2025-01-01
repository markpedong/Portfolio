package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func (h *handler) getWebsite(w http.ResponseWriter, r *http.Request) {
	t, err := h.client.GetWebsite(h.ctx, &pb.Empty{})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", models.Website{
		ID:        t.Id,
		Status:    int(t.Status),
		CreatedAt: t.CreatedAt.AsTime(),
		UpdatedAt: t.UpdatedAt.AsTime(),
		DeletedAt: utils.TimestampToTimePtr(t.DeletedAt),
	})
}

func (h *handler) updateWebsite(w http.ResponseWriter, r *http.Request) {
	var body models.Website
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}

	_, err := h.client.UpdateWebsite(h.ctx, &pb.WebsiteReq{
		Id:     body.ID,
		Status: int32(body.Status),
	})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
