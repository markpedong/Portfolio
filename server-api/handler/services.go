package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toServiceRes(links []*pb.ServiceRes) []*models.Services {
	var res []*models.Services
	for _, l := range links {
		res = append(res, &models.Services{
			ID:          l.Id,
			Title:       l.Title,
			Description: l.Description,
			Logo:        l.Logo,
			Status:      int(l.Status),
			CreatedAt:   l.CreatedAt.AsTime(),
			UpdatedAt:   l.UpdatedAt.AsTime(),
			DeletedAt:   utils.TimestampToTimePtr(l.DeletedAt),
		})
	}

	return res
}

func (h *handler) getServices(w http.ResponseWriter, r *http.Request) {
	services, err := h.client.GetServices(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toServiceRes(services.Services))
}

func (h *handler) createUpdateServices(w http.ResponseWriter, r *http.Request) {
	var body models.Services
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.ServiceRes{
		Id:          body.ID,
		Title:       body.Title,
		Description: body.Description,
		Logo:        body.Logo,
	}

	if fullPath[1] == "update" {
		payload.Id = body.ID
	}
	_, err := h.client.CreateUpdateServices(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
