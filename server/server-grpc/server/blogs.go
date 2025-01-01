package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateUpdateBlogs(ctx context.Context, b *pb.BlogsRes) (*pb.Empty, error) {
	if b.Id != "" {
		return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &models.Blogs{
			ID:          b.Id,
			Title:       b.Title,
			Description: b.Description,
			Image:       b.Image,
			Date:        b.Date,
			Link:        b.Link,
		}, utils.UpdateBlog)
	} else {
		return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Blogs{
			Title:       b.Title,
			Description: b.Description,
			Image:       b.Image,
			Date:        b.Date,
			Link:        b.Link,
		}, utils.CreateBlog)
	}
}

func (s *Server) GetBlogs(ctx context.Context, in *pb.Empty) (*pb.ListBlogsRes, error) {
	var l []models.Blogs
	err := s.storer.GetAllByModel(ctx, &l, "blogs", in.OnStatus)
	if err != nil {
		return nil, err
	}

	var pbBlogs []*pb.BlogsRes
	for _, w := range l {
		pbBlogs = append(pbBlogs, &pb.BlogsRes{
			Id:          w.ID,
			Title:       w.Title,
			Description: w.Description,
			Image:       w.Image,
			Date:        w.Date,
			Link:        w.Link,
			Status:      int32(w.Status),
			CreatedAt:   timestamppb.New(w.CreatedAt),
			UpdatedAt:   timestamppb.New(w.UpdatedAt),
			DeletedAt:   utils.DeletedAtNil(w.DeletedAt),
		})
	}

	return &pb.ListBlogsRes{Blogs: pbBlogs}, nil
}
