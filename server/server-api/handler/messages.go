package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toMessageRes(messages []*pb.MessageRes) []*models.Messages {
	var res []*models.Messages
	for _, v := range messages {
		res = append(res, &models.Messages{
			ID:        v.Id,
			Name:      v.Name,
			Email:     v.Email,
			Message:   v.Message,
			Status:    int(v.Status),
			CreatedAt: v.CreatedAt.AsTime(),
			UpdatedAt: v.UpdatedAt.AsTime(),
			DeletedAt: utils.TimestampToTimePtr(v.DeletedAt),
		})
	}
	return res
}

func (h *handler) getMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.client.GetMessages(h.ctx, &pb.Empty{})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toMessageRes(messages.Messages))
}

func (h *handler) sendMessage(w http.ResponseWriter, r *http.Request) {
	var body models.Messages
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}

	_, err := h.client.SendMessage(h.ctx, &pb.MessageRes{
		Name:    body.Name,
		Email:   body.Email,
		Message: body.Message,
	})
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
