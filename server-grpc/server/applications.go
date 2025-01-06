package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateUpdateApplications(ctx context.Context, a *pb.AppRes) (*pb.Empty, error) {
	if a.Id != "" {
		return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &models.Application{
			ID:    a.Id,
			Image: a.Image,
			Name:  a.Name,
		}, utils.UpdateApplication)
	} else {
		return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Application{
			Image: a.Image,
			Name:  a.Name,
		}, utils.CreateApp)
	}
}

func (s *Server) GetApplications(ctx context.Context, in *pb.Empty) (*pb.ListAppRes, error) {
	var w []models.Application
	err := s.storer.GetAllByModel(ctx, &w, "applications", in.OnStatus)
	if err != nil {
		return nil, err
	}

	var pbApps []*pb.AppRes
	for _, l := range w {
		pbApps = append(pbApps, &pb.AppRes{
			Id:        l.ID,
			Name:      l.Name,
			Image:     l.Image,
			Status:    int32(l.Status),
			CreatedAt: timestamppb.New(l.CreatedAt),
			UpdatedAt: timestamppb.New(l.UpdatedAt),
			DeletedAt: utils.DeletedAtNil(l.DeletedAt),
		})
	}

	return &pb.ListAppRes{Applications: pbApps}, nil
}
