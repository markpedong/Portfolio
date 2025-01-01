package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPbLinks(l []models.Links) *pb.ListLinkRes {
	var res []*pb.LinkRes
	for _, v := range l {
		res = append(res, &pb.LinkRes{
			Id:        v.ID,
			Link:      v.Link,
			Type:      v.Type,
			Status:    int32(v.Status),
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
			DeletedAt: utils.DeletedAtNil(v.DeletedAt),
		})
	}

	return &pb.ListLinkRes{Links: res}
}

func (s *Server) GetLinks(ctx context.Context, in *pb.Empty) (*pb.ListLinkRes, error) {
	var l []models.Links
	err := s.storer.GetAllByModel(ctx, &l, "links", in.OnStatus)
	if err != nil {
		return nil, err
	}

	return toPbLinks(l), nil
}
func (s *Server) CreateUpdateLinks(ctx context.Context, e *pb.LinkRes) (*pb.Empty, error) {
	if e.Id != "" {
		return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &models.Links{
			ID:   e.Id,
			Link: e.Link,
			Type: e.Type,
		}, utils.UpdateLink)
	}

	return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Links{
		Link: e.Link,
		Type: e.Type,
	}, utils.CreateLink)
}
