package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
)

func (s *Server) GetEducations(ctx context.Context, in *pb.Empty) (*pb.ListEduRes, error) {
	var edu []models.Education
	err := s.storer.GetAllByModel(ctx, &edu, "educations", in.OnStatus)
	if err != nil {
		return nil, err
	}

	for i := range edu {
		var skills []models.EduSkill
		err = s.storer.GetByModel(ctx, &skills, edu[i].ID, "edu_skill", "education_id")
		if err != nil {
			return nil, err
		}
		edu[i].Skills = skills
	}

	return toPbEducationRes(edu), nil
}

func (s *Server) GetEducation(ctx context.Context, e *pb.IDReq) (*pb.EduRes, error) {
	var edu models.Education
	err := s.storer.GetByModel(ctx, &edu, e.Id, "educations")
	if err != nil {
		return nil, err
	}

	var skills []models.EduSkill
	err = s.storer.GetByModel(ctx, &skills, edu.ID, "edu_skill", "education_id")
	if err != nil {
		return nil, err
	}
	edu.Skills = skills

	return toPbEducation(&edu), nil
}

func (s *Server) CreateUpdateEducations(ctx context.Context, e *pb.EduRes) (*pb.Empty, error) {
	if e.Id != "" {
		var skills []models.EduSkill
		for _, w := range e.Skills {
			skills = append(skills, models.EduSkill{
				Name:       w.Name,
				Percentage: int(w.Percentage),
			})
		}

		_, err := s.storer.UpdateEducation(ctx, &models.Education{
			ID:          e.Id,
			School:      e.School,
			Course:      e.Course,
			Started:     e.Started,
			Ended:       e.Ended,
			Description: e.Description,
			Skills:      skills,
		})
		if err != nil {
			return nil, err
		}

		return &pb.Empty{}, nil
	} else {
		return &pb.Empty{}, s.storer.CreateEducation(ctx, toStorerEducation(e))
	}
}
