package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toStorerEduSkill(e []models.EduSkill) []*pb.EduSkillRes {
	var es []*pb.EduSkillRes

	for _, w := range e {
		es = append(es, &pb.EduSkillRes{
			Name:       w.Name,
			Percentage: int32(w.Percentage),
		})
	}

	return es
}

func toEduSkillRes(e []*pb.EduSkillRes) []models.EduSkill {
	es := make([]models.EduSkill, 0, len(e))
	for _, w := range e {
		es = append(es, models.EduSkill{
			ID:          w.Id,
			Name:        w.Name,
			Percentage:  int(w.Percentage),
			EducationID: w.EducationId,
		})
	}
	return es
}

func toEducationRes(edu []*pb.EduRes) []*models.Education {
	var es []*models.Education

	for _, w := range edu {
		es = append(es, &models.Education{
			ID:          w.Id,
			School:      w.School,
			Course:      w.Course,
			Started:     w.Started,
			Ended:       w.Ended,
			Description: w.Description,
			Skills:      toEduSkillRes(w.Skills),
			Status:      int(w.Status),
			CreatedAt:   w.CreatedAt.AsTime(),
			UpdatedAt:   w.UpdatedAt.AsTime(),
			DeletedAt:   utils.TimestampToTimePtr(w.DeletedAt),
		})
	}

	return es
}

// /education
func (h *handler) createUpdateEdu(w http.ResponseWriter, r *http.Request) {
	var e models.Education
	if err := helpers.BindValidateJSON(w, r, &e); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.EduRes{
		School:      e.School,
		Course:      e.Course,
		Started:     e.Started,
		Ended:       e.Ended,
		Description: e.Description,
		Skills:      toStorerEduSkill(e.Skills),
	}

	if fullPath[1] == "update" {
		payload.Id = e.ID
	}
	_, err := h.client.CreateUpdateEducations(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}

func (h *handler) getEducations(w http.ResponseWriter, r *http.Request) {
	educations, err := h.client.GetEducations(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toEducationRes(educations.Educations))
}
