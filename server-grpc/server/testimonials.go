package server

import (
	"context"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPbTestimonials(q []*models.Testimonials) *pb.ListTestimonialRes {
	var pbTestimonials []*pb.TestimonialRes
	for _, l := range q {
		pbTestimonials = append(pbTestimonials, &pb.TestimonialRes{
			Id:          l.ID,
			Author:      l.Author,
			Description: l.Description,
			Job:         l.Job,
			Image:       l.Image,
			Status:      int32(l.Status),
			CreatedAt:   timestamppb.New(l.CreatedAt),
			UpdatedAt:   timestamppb.New(l.UpdatedAt),
			DeletedAt:   utils.DeletedAtNil(l.DeletedAt),
		})
	}

	return &pb.ListTestimonialRes{Testimonials: pbTestimonials}
}

func (s *Server) GetTestimonials(ctx context.Context, in *pb.Empty) (*pb.ListTestimonialRes, error) {
	var l []*models.Testimonials
	err := s.storer.GetAllByModel(ctx, &l, "testimonials", in.OnStatus)
	if err != nil {
		return nil, err
	}

	return toPbTestimonials(l), nil
}

func (s *Server) CreateUpdateTestimonials(ctx context.Context, req *pb.TestimonialRes) (*pb.Empty, error) {
	if req.Id != "" {
		return &pb.Empty{}, s.storer.UpdateRowByModel(ctx, &models.Testimonials{
			ID:          req.Id,
			Author:      req.Author,
			Image:       req.Image,
			Job:         req.Job,
			Description: req.Description,
		},
			utils.UpdateTestimonial,
		)
	} else {
		return &pb.Empty{}, s.storer.CreateRowByModel(ctx, &models.Testimonials{
			Author:      req.Author,
			Description: req.Description,
			Image:       req.Image,
			Job:         req.Job,
		},
			utils.CreateTestimonial,
		)
	}
}
