package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toAppsRes(apps []*pb.AppRes) []*models.Application {
	var res []*models.Application
	for _, app := range apps {
		res = append(res, &models.Application{
			ID:        app.Id,
			Image:     app.Image,
			Name:      app.Name,
			Status:    int(app.Status),
			CreatedAt: app.CreatedAt.AsTime(),
			UpdatedAt: app.UpdatedAt.AsTime(),
			DeletedAt: utils.TimestampToTimePtr(app.DeletedAt),
		})
	}

	return res
}

func (h *handler) getApplications(w http.ResponseWriter, r *http.Request) {
	apps, err := h.client.GetApplications(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toAppsRes(apps.Applications))
}

func (h *handler) createUpdateApplications(w http.ResponseWriter, r *http.Request) {
	var body models.Application
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.AppRes{
		Name:  body.Name,
		Image: body.Image,
	}

	if fullPath[1] == "update" {
		payload.Id = body.ID
	}
	_, err := h.client.CreateUpdateApplications(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
