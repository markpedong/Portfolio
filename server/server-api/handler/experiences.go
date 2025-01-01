package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
)

func (h *handler) getExperiences(w http.ResponseWriter, r *http.Request) {
	experiences, err := h.client.GetExperiences(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []models.Experiences
	for _, exp := range experiences.Experiences {
		var s []models.ExpSkill
		for _, v := range exp.Skills {
			s = append(s, models.ExpSkill{
				Name:         v.Name,
				Percentage:   int(v.Percentage),
				ID:           v.Id,
				ExperienceID: v.ExperienceId,
			})
		}
		resp = append(resp, models.Experiences{
			ID:           exp.Id,
			Company:      exp.Company,
			Title:        exp.Title,
			Location:     exp.Location,
			Descriptions: exp.Descriptions,
			Skills:       s,
			Started:      exp.Started,
			Ended:        exp.Ended,
			Status:       int(exp.Status),
			CreatedAt:    exp.CreatedAt.AsTime(),
			UpdatedAt:    exp.UpdatedAt.AsTime(),
		})
	}

	helpers.JSONResponse(w, "", resp)
}

func (h *handler) createUpdateExp(w http.ResponseWriter, r *http.Request) {
	var body models.Experiences
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}

	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.ExpRes{
		Company:      body.Company,
		Title:        body.Title,
		Location:     body.Location,
		Descriptions: body.Descriptions,
		Started:      body.Started,
		Ended:        body.Ended,
		Status:       int32(body.Status),
		Skills:       toPBExpSkillReq(&body.Skills),
	}
	if fullPath[1] == "update" {
		payload.Id = body.ID
	}

	_, err := h.client.CreateUpdateExperiences(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}

func toPBExpSkillReq(e *[]models.ExpSkill) []*pb.ExpSkillRes {
	var items []*pb.ExpSkillRes
	for _, item := range *e {
		items = append(items, &pb.ExpSkillRes{
			Name:       item.Name,
			Percentage: int32(item.Percentage),
		})
	}

	return items
}
