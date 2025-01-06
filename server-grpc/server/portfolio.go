package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPbPortfolios(p []*models.Portfolios) *pb.ListPortfolioRes {
	var res []*pb.PortfolioRes
	for _, v := range p {
		res = append(res, &pb.PortfolioRes{
			Id:        v.ID,
			Title:     v.Title,
			Tech:      v.Tech,
			Link:      v.Link,
			Image:     v.Image,
			Status:    int32(v.Status),
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
			DeletedAt: utils.DeletedAtNil(v.DeletedAt),
		})
	}

	return &pb.ListPortfolioRes{
		Portfolios: res,
	}
}

func (s *Server) GetPortfolios(ctx context.Context, in *pb.Empty) (*pb.ListPortfolioRes, error) {
	var p []*models.Portfolios
	if err := s.storer.GetAllByModel(ctx, &p, "portfolios", in.OnStatus); err != nil {
		return nil, err
	}

	return toPbPortfolios(p), nil
}

func (s *Server) CreateUpdatePortfolios(ctx context.Context, req *pb.PortfolioRes) (*pb.Empty, error) {
	if req.Id != "" {
		return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &models.Portfolios{
			ID:    req.Id,
			Title: req.Title,
			Link:  req.Link,
			Tech:  req.Tech,
			Image: req.Image,
		}, utils.UpdatePortfolio)
	} else {
		return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Portfolios{
			Title: req.Title,
			Link:  req.Link,
			Tech:  req.Tech,
			Image: req.Image,
		}, utils.CreatePortfolio)
	}
}
