package server

import (
	"context"
	"fmt"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/server-grpc/storer"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	storer *storer.PSQLStorer
	pb.UnimplementedApiServiceServer
}

func NewServer(storer *storer.PSQLStorer) *Server {
	return &Server{
		storer: storer,
	}
}

func (s *Server) GetUser(ctx context.Context, u *pb.UserRes) (*pb.UserRes, error) {
	var user models.Users
	err := s.storer.GetByModel(ctx, &user, u.Username, "users", "username")
	if err != nil {
		return nil, err
	}

	return toUserRes(&user), nil
}

func (s *Server) CreateUpdateUser(ctx context.Context, u *pb.UserRes) (*pb.Empty, error) {
	if u.Id == "" {
		return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Users{
			FirstName:      u.FirstName,
			LastName:       u.LastName,
			Phone:          u.Phone,
			Address:        u.Address,
			Description:    u.Description,
			Email:          u.Email,
			Username:       u.Username,
			Password:       u.Password,
			ResumePDF:      u.ResumePdf,
			ResumeDocx:     u.ResumeDocx,
			IsDownloadable: int(u.Isdownloadable),
		}, utils.CreateUser)
	}

	return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &models.Users{
		ID:             u.Id,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Phone:          u.Phone,
		Address:        u.Address,
		Description:    u.Description,
		Email:          u.Email,
		Username:       u.Username,
		Password:       u.Password,
		ResumePDF:      u.ResumePdf,
		ResumeDocx:     u.ResumeDocx,
		IsDownloadable: int(u.Isdownloadable),
	}, utils.UpdateUser)
}

func (s *Server) CreateUpdateSession(ctx context.Context, su *pb.SessionRes) (*pb.SessionRes, error) {
	if su.CreatedAt != nil {
		gs, err := s.storer.UpdateSession(ctx, toStorerSession(su))
		if err != nil {
			return nil, err
		}

		return toPbSession(gs), nil
	} else {
		gs, err := s.storer.CreateSession(ctx, toStorerSession(su))
		if err != nil {
			return nil, err
		}

		return toPbSession(gs), nil
	}
}

func (s *Server) GetSession(ctx context.Context, id *pb.IDReq) (*pb.SessionRes, error) {
	var gs models.Session
	err := s.storer.GetByModel(ctx, &gs, id.Id, "sessions")
	if err != nil {
		return nil, err
	}

	return toPbSession(&gs), nil
}

func (s *Server) GetSessions(ctx context.Context, _ *pb.Empty) (*pb.ListSessionsRes, error) {
	var gs []models.Session
	err := s.storer.GetAllByModel(ctx, &gs, "sessions")
	if err != nil {
		return nil, err
	}

	return toPbSessions(gs), nil
}

func (s *Server) ToggleOrDelete(ctx context.Context, body *pb.IdModel) (*pb.Empty, error) {
	return &pb.Empty{}, s.storer.ToggleOrDelete(ctx, body)
}

func (s *Server) GetWebsite(ctx context.Context, _ *pb.Empty) (*pb.WebsiteRes, error) {
	var websites []models.Website
	err := s.storer.GetAllByModel(ctx, &websites, "website")
	if err != nil {
		return nil, err
	}

	if len(websites) == 0 {
		return &pb.WebsiteRes{}, nil
	}

	return &pb.WebsiteRes{Id: websites[0].ID,
		Status:    int32(websites[0].Status),
		CreatedAt: timestamppb.New(websites[0].CreatedAt),
		UpdatedAt: timestamppb.New(websites[0].UpdatedAt),
		DeletedAt: utils.DeletedAtNil(websites[0].DeletedAt),
	}, nil
}

func (s *Server) UpdateWebsite(ctx context.Context, w *pb.WebsiteReq) (*pb.Empty, error) {
	return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &pb.WebsiteReq{
		Id:     w.Id,
		Status: w.Status,
	}, utils.UpdateWebsite)
}

func (s *Server) GetPublicDetails(ctx context.Context, _ *pb.Empty) (*pb.UserRes, error) {
	var user []models.Users
	err := s.storer.GetAllByModel(ctx, &user, "users")
	if err != nil {
		return nil, err
	}

	if len(user) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return toPbUser(&user[0]), nil
}

func toPbUser(u *models.Users) *pb.UserRes {
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
	}
}
