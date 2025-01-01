package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) SendMessage(ctx context.Context, req *pb.MessageRes) (*pb.Empty, error) {
	return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Messages{
		Name:    req.Name,
		Email:   req.Email,
		Message: req.Message,
	}, utils.CreateMessage)
}

func (s *Server) GetMessages(ctx context.Context, req *pb.Empty) (*pb.ListMessageRes, error) {
	var messages []*models.Messages
	err := s.storer.GetAllByModel(ctx, &messages, "messages")
	if err != nil {
		return nil, err
	}

	var Messages []*pb.MessageRes
	for _, v := range messages {
		Messages = append(Messages, &pb.MessageRes{
			Id:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Message:   v.Message,
			Status:    int32(v.Status),
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
			DeletedAt: utils.DeletedAtNil(v.DeletedAt),
		})
	}

	return &pb.ListMessageRes{Messages: Messages}, nil
}
