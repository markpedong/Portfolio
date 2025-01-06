package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toLinksRes(links []*pb.LinkRes) []*models.Links {
	var res []*models.Links
	for _, l := range links {
		res = append(res, &models.Links{
			ID:        l.Id,
			Link:      l.Link,
			Type:      l.Type,
			Status:    int(l.Status),
			CreatedAt: l.CreatedAt.AsTime(),
			UpdatedAt: l.UpdatedAt.AsTime(),
			DeletedAt: utils.TimestampToTimePtr(l.DeletedAt),
		})
	}

	return res
}

func (h *handler) getLinks(w http.ResponseWriter, r *http.Request) {
	links, err := h.client.GetLinks(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toLinksRes(links.Links))
}

func (h *handler) createUpdateLinks(w http.ResponseWriter, r *http.Request) {
	var body models.Links
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.LinkRes{
		Link: body.Link,
		Type: body.Type,
	}

	if fullPath[1] == "update" {
		payload.Id = body.ID
	}
	_, err := h.client.CreateUpdateLinks(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
