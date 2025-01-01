package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toTestimonialRes(t []*pb.TestimonialRes) []*models.Testimonials {
	var res []*models.Testimonials
	for _, l := range t {
		res = append(res, &models.Testimonials{
			ID:          l.Id,
			Author:      l.Author,
			Description: l.Description,
			Image:       l.Image,
			Job:         l.Job,
			Status:      int(l.Status),
			CreatedAt:   l.CreatedAt.AsTime(),
			UpdatedAt:   l.UpdatedAt.AsTime(),
			DeletedAt:   utils.TimestampToTimePtr(l.DeletedAt),
		})
	}

	return res
}

func (h *handler) getTestimonials(w http.ResponseWriter, r *http.Request) {
	testimonials, err := h.client.GetTestimonials(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toTestimonialRes(testimonials.Testimonials))
}

func (h *handler) createUpdateTestimonials(w http.ResponseWriter, r *http.Request) {
	var body models.Testimonials
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.TestimonialRes{
		Author:      body.Author,
		Description: body.Description,
		Image:       body.Image,
		Job:         body.Job,
	}

	if fullPath[1] == "update" {
		payload.Id = body.ID
	}
	_, err := h.client.CreateUpdateTestimonials(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
