package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetFiles(ctx context.Context, _ *pb.Empty) (*pb.ListFileRes, error) {
	var files []models.Files
	err := s.storer.GetAllByModel(ctx, &files, "files")
	if err != nil {
		return nil, err
	}

	var pbFiles []*pb.FileRes
	for _, file := range files {
		pbFiles = append(pbFiles, &pb.FileRes{
			Id:        file.ID,
			Name:      file.Name,
			File:      file.File,
			CreatedAt: timestamppb.New(file.CreatedAt),
		})
	}

	return &pb.ListFileRes{
		Files: pbFiles,
	}, nil
}

func (s *Server) CreateFile(ctx context.Context, req *pb.FileReq) (*pb.Empty, error) {
	return &pb.Empty{}, s.storer.CreateFile(ctx, &models.Files{
		File: req.File,
		Name: req.Name,
	})
}
