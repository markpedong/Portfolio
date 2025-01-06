package handler

import (
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"portfolio/utils"
)

func toBlogRes(blogs []*pb.BlogsRes) []*models.Blogs {
	var res []*models.Blogs
	for _, blog := range blogs {
		res = append(res, &models.Blogs{
			ID:          blog.Id,
			Title:       blog.Title,
			Date:        blog.Date,
			Description: blog.Description,
			Link:        blog.Link,
			Image:       blog.Image,
			Status:      int(blog.Status),
			CreatedAt:   blog.CreatedAt.AsTime(),
			UpdatedAt:   blog.UpdatedAt.AsTime(),
			DeletedAt:   utils.TimestampToTimePtr(blog.DeletedAt),
		})
	}

	return res
}

func (h *handler) getBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.client.GetBlogs(h.ctx, h.isPulicRoute(r))
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "", toBlogRes(blogs.Blogs))
}

func (h *handler) createUpdateBlogs(w http.ResponseWriter, r *http.Request) {
	var body models.Blogs
	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
		return
	}
	fullPath := cleanSplit(r.Context().Value(fullPath{}).(string))
	payload := pb.BlogsRes{
		Title:       body.Title,
		Date:        body.Date,
		Description: body.Description,
		Link:        body.Link,
		Image:       body.Image,
	}

	if fullPath[1] == "update" {
		payload.Id = body.ID
	}
	_, err := h.client.CreateUpdateBlogs(h.ctx, &payload)
	if err != nil {
		helpers.ErrJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, "")
}
