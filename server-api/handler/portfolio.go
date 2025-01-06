package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toPortfolio(p *pb.PortfolioRes) *models.Portfolios {
	return &models.Portfolios{
		ID:        p.Id,
		Title:     p.Title,
		Tech:      p.Tech,
		Link:      p.Link,
		Image:     p.Image,
		Status:    int(p.Status),
		CreatedAt: p.CreatedAt.AsTime(),
		UpdatedAt: p.UpdatedAt.AsTime(),
		DeletedAt: utils.TimestampToTimePtr(p.DeletedAt),
	}
}
func toPortfolioRes(portfolios []*pb.PortfolioRes) []*models.Portfolios {
	var res []*models.Portfolios
	for _, p := range portfolios {
		res = append(res, toPortfolio(p))
	}

	return res
}

func (h *handler) getPortfolios(w http.ResponseWriter, r *http.Request) {
	portfolio, err := h.client.GetPortfolios(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toPortfolioRes(portfolio.Portfolios))
}

func (h *handler) createUpdatePortfolios(w http.ResponseWriter, r *http.Request) {
	var body models.Portfolios
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.PortfolioRes{
		Title: body.Title,
		Tech:  body.Tech,
		Link:  body.Link,
		Image: body.Image,
	}
	if fullPath[1] == "update" {
		payload.Id = body.ID
	}

	_, err := h.client.CreateUpdatePortfolios(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
