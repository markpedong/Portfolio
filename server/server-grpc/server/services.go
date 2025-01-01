package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPbServices(l []*models.Services) *pb.ListServiceRes {
	var res []*pb.ServiceRes
	for _, v := range l {
		res = append(res, &pb.ServiceRes{
			Id:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Logo:        v.Logo,
			Status:      int32(v.Status),
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
			DeletedAt:   utils.DeletedAtNil(v.DeletedAt),
		})
	}

	return &pb.ListServiceRes{Services: res}
}

func (s *Server) GetServices(ctx context.Context, in *pb.Empty) (*pb.ListServiceRes, error) {
	var l []*models.Services
	err := s.storer.GetAllByModel(ctx, &l, "services", in.OnStatus)
	if err != nil {
		return nil, err
	}

	return toPbServices(l), nil
}

func (s *Server) CreateUpdateServices(ctx context.Context, req *pb.ServiceRes) (*pb.Empty, error) {
	if req.Id != "" {
		return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &models.Services{
			ID:          req.Id,
			Title:       req.Title,
			Description: req.Description,
			Logo:        req.Logo,
		},
			utils.UpdateService,
		)
	} else {
		return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Services{
			Title:       req.Title,
			Description: req.Description,
			Logo:        req.Logo,
		},
			utils.CreateService,
		)
	}
}
