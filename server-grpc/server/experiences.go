package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
)

func (s *Server) GetExperiences(ctx context.Context, in *pb.Empty) (*pb.ListExpRes, error) {
	var exp []models.Experiences
	err := s.storer.GetAllByModel(ctx, &exp, "experiences", in.OnStatus)
	if err != nil {
		return nil, err
	}

	for i := range exp {
		var es []models.ExpSkill
		err = s.storer.GetByModel(ctx, &es, exp[i].ID, "exp_skill", "experience_id")
		if err != nil {
			return nil, err
		}

		exp[i].Skills = es
	}

	return toPbServicesRes(exp), nil
}

func (s *Server) CreateUpdateExperiences(ctx context.Context, req *pb.ExpRes) (*pb.Empty, error) {
	if req.Id == "" {
		return &pb.Empty{}, s.storer.CreateExperience(ctx, toExperienceReq(req))
	} else {
		var skills []models.ExpSkill
		for _, v := range req.Skills {
			skills = append(skills, models.ExpSkill{
				Name:         v.Name,
				Percentage:   int(v.Percentage),
				ExperienceID: v.ExperienceId,
			})
		}

		_, err := s.storer.UpdateExperience(ctx, &models.Experiences{
			ID:           req.Id,
			Company:      req.Company,
			Title:        req.Title,
			Location:     req.Location,
			Started:      req.Started,
			Ended:        req.Ended,
			Descriptions: req.Descriptions,
			Skills:       skills,
		})
		if err != nil {
			return nil, err
		}

		return &pb.Empty{}, nil
	}
}

func (s *Server) GetExperience(ctx context.Context, e *pb.IDReq) (*pb.ExpRes, error) {
	var exp models.Experiences
	err := s.storer.GetByModel(ctx, &exp, e.Id, "experiences")
	if err != nil {
		return nil, err
	}

	var skills []models.ExpSkill
	err = s.storer.GetByModel(ctx, &skills, exp.ID, "exp_skill", "experience_id")
	if err != nil {
		return nil, err
	}
	exp.Skills = skills

	return toPbExperience(&exp), nil
}
