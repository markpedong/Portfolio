package server

import (
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toUserRes(u *models.Users) *pb.UserRes {
	return &pb.UserRes{
		Id:             u.ID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Phone:          u.Phone,
		Address:        u.Address,
		Description:    u.Description,
		Email:          u.Email,
		Username:       u.Username,
		Password:       u.Password,
		ResumePdf:      u.ResumePDF,
		ResumeDocx:     u.ResumeDocx,
		Isdownloadable: int32(u.IsDownloadable),
		CreatedAt:      timestamppb.New(u.CreatedAt),
		UpdatedAt:      timestamppb.New(u.UpdatedAt),
		DeletedAt:      utils.DeletedAtNil(u.DeletedAt),
	}
}

func toStorerSession(s *pb.SessionRes) *models.Session {
	return &models.Session{
		ID:           s.Id,
		UserID:       s.UserId,
		Email:        s.Email,
		RefreshToken: s.RefreshToken,
		IsRevoked:    s.IsRevoked,
		CreatedAt:    s.CreatedAt.AsTime(),
		ExpiresAt:    utils.TimestampToTimePtr(s.ExpiresAt),
	}
}

func toPbSession(s *models.Session) *pb.SessionRes {
	return &pb.SessionRes{
		Id:           s.ID,
		UserId:       s.UserID,
		Email:        s.Email,
		RefreshToken: s.RefreshToken,
		IsRevoked:    s.IsRevoked,
		CreatedAt:    timestamppb.New(s.CreatedAt),
		ExpiresAt:    timestamppb.New(*s.ExpiresAt),
	}
}

func toPbSessions(s []models.Session) *pb.ListSessionsRes {
	pbSessions := &pb.ListSessionsRes{}
	for _, session := range s {
		pbSessions.Sessions = append(pbSessions.Sessions, toPbSession(&session))
	}

	return pbSessions
}

func toPbServicesRes(experiences []models.Experiences) *pb.ListExpRes {
	var items []*pb.ExpRes
	for _, exp := range experiences {
		var s []*pb.ExpSkillRes

		for _, v := range exp.Skills {
			s = append(s, &pb.ExpSkillRes{
				Id:           v.ID,
				Name:         v.Name,
				Percentage:   int32(v.Percentage),
				ExperienceId: v.ExperienceID,
			})
		}
		items = append(items, &pb.ExpRes{
			Id:           exp.ID,
			Company:      exp.Company,
			Title:        exp.Title,
			Location:     exp.Location,
			Status:       int32(exp.Status),
			Descriptions: exp.Descriptions,
			Started:      exp.Started,
			Ended:        exp.Ended,
			Skills:       s,
			CreatedAt:    timestamppb.New(exp.CreatedAt),
			UpdatedAt:    timestamppb.New(exp.UpdatedAt),
			DeletedAt:    utils.DeletedAtNil(exp.DeletedAt),
		})
	}

	return &pb.ListExpRes{
		Experiences: items,
	}
}
func toExperienceReq(e *pb.ExpRes) *models.Experiences {
	var s []models.ExpSkill
	for _, v := range e.Skills {
		s = append(s, models.ExpSkill{
			Name:       v.Name,
			Percentage: int(v.Percentage),
		})
	}

	return &models.Experiences{
		Company:      e.Company,
		Title:        e.Title,
		Location:     e.Location,
		Started:      e.Started,
		Ended:        e.Ended,
		Descriptions: e.Descriptions,
		Skills:       s,
	}
}

func toPbExperience(e *models.Experiences) *pb.ExpRes {
	var s []*pb.ExpSkillRes
	for _, v := range e.Skills {
		s = append(s, &pb.ExpSkillRes{
			Name:         v.Name,
			Percentage:   int32(v.Percentage),
			ExperienceId: v.ExperienceID,
			Id:           v.ID,
		})
	}
	return &pb.ExpRes{
		Id:           e.ID,
		Company:      e.Company,
		Title:        e.Title,
		Location:     e.Location,
		Descriptions: e.Descriptions,
		Started:      e.Started,
		Ended:        e.Ended,
		Skills:       s,
		Status:       int32(e.Status),
		CreatedAt:    timestamppb.New(e.CreatedAt),
		UpdatedAt:    timestamppb.New(e.UpdatedAt),
		DeletedAt:    utils.DeletedAtNil(e.DeletedAt),
	}
}

func toStorerEduSkill(e []*pb.EduSkillRes) []models.EduSkill {
	var es []models.EduSkill

	for _, w := range e {
		es = append(es, models.EduSkill{
			Name:       w.Name,
			Percentage: int(w.Percentage),
		})
	}

	return es
}

func toStorerEducation(e *pb.EduRes) *models.Education {
	return &models.Education{
		School:      e.School,
		Course:      e.Course,
		Started:     e.Started,
		Ended:       e.Ended,
		Description: e.Description,
		Skills:      toStorerEduSkill(e.Skills),
	}
}

func toPbEduSkillRes(es []models.EduSkill) []*pb.EduSkillRes {
	pbEduSkillRes := []*pb.EduSkillRes{}
	for _, e := range es {
		pbEduSkillRes = append(pbEduSkillRes, &pb.EduSkillRes{
			Id:          e.ID,
			Name:        e.Name,
			Percentage:  int32(e.Percentage),
			EducationId: e.EducationID,
		})
	}

	return pbEduSkillRes
}
func toPbEducation(e *models.Education) *pb.EduRes {
	return &pb.EduRes{
		Id:          e.ID,
		School:      e.School,
		Course:      e.Course,
		Started:     e.Started,
		Ended:       e.Ended,
		Description: e.Description,
		Status:      int32(e.Status),
		Skills:      toPbEduSkillRes(e.Skills),
		CreatedAt:   timestamppb.New(e.CreatedAt),
		UpdatedAt:   timestamppb.New(e.UpdatedAt),
		DeletedAt:   utils.DeletedAtNil(e.DeletedAt),
	}
}

func toPbEducationRes(edu []models.Education) *pb.ListEduRes {
	pbEducations := &pb.ListEduRes{}
	for _, education := range edu {
		pbEducations.Educations = append(pbEducations.Educations, toPbEducation(&education))
	}

	return pbEducations
}
