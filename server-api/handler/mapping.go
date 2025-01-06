package handler

import (
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toUserResPb(gu *pb.UserRes) models.Users {
	return models.Users{
		ID:             gu.Id,
		FirstName:      gu.FirstName,
		LastName:       gu.LastName,
		Phone:          gu.Phone,
		Address:        gu.Address,
		Description:    gu.Description,
		Email:          gu.Email,
		Username:       gu.Username,
		Password:       gu.Password,
		ResumePDF:      gu.ResumePdf,
		ResumeDocx:     gu.ResumeDocx,
		IsDownloadable: int(gu.Isdownloadable),
		CreatedAt:      gu.CreatedAt.AsTime(),
		UpdatedAt:      gu.UpdatedAt.AsTime(),
		DeletedAt:      utils.TimestampToTimePtr(gu.DeletedAt),
	}
}
