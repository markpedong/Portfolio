package handler

import (
	"fmt"
	"log"
	"net/http"
	"portfolio/models"
)

var r *http.ServeMux

func CreateRoutes(handler *handler) {
	r = http.NewServeMux()
	tokenMaker := handler.tokenMaker
	public := http.NewServeMux()
	applications := http.NewServeMux()
	blogs := http.NewServeMux()
	files := http.NewServeMux()
	educations := http.NewServeMux()
	experiences := http.NewServeMux()
	links := http.NewServeMux()
	messages := http.NewServeMux()
	services := http.NewServeMux()
	users := http.NewServeMux()
	tokens := http.NewServeMux()
	portfolios := http.NewServeMux()
	logs := http.NewServeMux()
	testimonials := http.NewServeMux()
	website := http.NewServeMux()

	// /applications
	applications.HandleFunc("GET /get", handler.getApplications)
	applications.HandleFunc("POST /add", handler.createUpdateApplications)
	applications.HandleFunc("POST /update", handler.createUpdateApplications)
	applications.HandleFunc("POST /delete", handler.toggleOrDelete)
	applications.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /blogs
	blogs.HandleFunc("GET /get", handler.getBlogs)
	blogs.HandleFunc("POST /add", handler.createUpdateBlogs)
	blogs.HandleFunc("POST /update", handler.createUpdateBlogs)
	blogs.HandleFunc("POST /delete", handler.toggleOrDelete)
	blogs.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /educations
	educations.HandleFunc("GET /get", handler.getEducations)
	educations.HandleFunc("POST /add", handler.createUpdateEdu)
	educations.HandleFunc("POST /update", handler.createUpdateEdu)
	educations.HandleFunc("POST /delete", handler.toggleOrDelete)
	educations.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /experiences
	experiences.HandleFunc("GET /get", handler.getExperiences)
	experiences.HandleFunc("POST /add", handler.createUpdateExp)
	experiences.HandleFunc("POST /update", handler.createUpdateExp)
	experiences.HandleFunc("POST /delete", handler.toggleOrDelete)
	experiences.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /links
	links.HandleFunc("GET /get", handler.getLinks)
	links.HandleFunc("POST /add", handler.createUpdateLinks)
	links.HandleFunc("POST /update", handler.createUpdateLinks)
	links.HandleFunc("POST /delete", handler.toggleOrDelete)
	links.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /files
	files.HandleFunc("GET /get", handler.getFiles)
	files.HandleFunc("POST /upload", handler.uploadImage)
	files.HandleFunc("POST /delete", handler.deleteFiles)

	// /logs
	logs.HandleFunc("GET /get", handler.getSession)

	// /messages
	messages.HandleFunc("GET /get", handler.getMessages)

	// /public
	public.HandleFunc("POST /login", handler.loginUser)
	public.HandleFunc("GET /details", handler.getPublicDetails)
	public.HandleFunc("GET /website", handler.getWebsite)
	public.HandleFunc("POST /sendMsg", handler.sendMessage)
	public.HandleFunc("GET /applications", handler.getApplications)
	public.HandleFunc("GET /blogs", handler.getBlogs)
	public.HandleFunc("GET /experiences", handler.getExperiences)
	public.HandleFunc("GET /educations", handler.getEducations)
	public.HandleFunc("GET /links", handler.getLinks)
	public.HandleFunc("GET /services", handler.getServices)
	public.HandleFunc("GET /portfolios", handler.getPortfolios)
	public.HandleFunc("GET /testimonials", handler.getTestimonials)
	// public.HandleFunc("POST /logout", handler.logoutUser)

	// /portfolios
	portfolios.HandleFunc("GET /get", handler.getPortfolios)
	portfolios.HandleFunc("POST /add", handler.createUpdatePortfolios)
	portfolios.HandleFunc("POST /update", handler.createUpdatePortfolios)
	portfolios.HandleFunc("POST /delete", handler.toggleOrDelete)
	portfolios.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /services
	services.HandleFunc("GET /get", handler.getServices)
	services.HandleFunc("POST /add", handler.createUpdateServices)
	services.HandleFunc("POST /update", handler.createUpdateServices)
	services.HandleFunc("POST /delete", handler.toggleOrDelete)
	services.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /testimonials
	testimonials.HandleFunc("GET /get", handler.getTestimonials)
	testimonials.HandleFunc("POST /add", handler.createUpdateTestimonials)
	testimonials.HandleFunc("POST /update", handler.createUpdateTestimonials)
	testimonials.HandleFunc("POST /delete", handler.toggleOrDelete)
	testimonials.HandleFunc("POST /toggle", handler.toggleOrDelete)

	// /users
	users.HandleFunc("POST /add", handler.createUpdateUser)
	users.HandleFunc("GET /get", handler.getUser)
	users.HandleFunc("POST /update", handler.createUpdateUser)

	// /tokens
	tokens.HandleFunc("POST /renew", handler.renewAccessToken)
	// tokens.HandleFunc("POST /revoke", handler.revokeSession)

	// /website
	website.HandleFunc("GET /get", handler.getWebsite)
	website.HandleFunc("POST /update", handler.updateWebsite)

	r.Handle("/public/", http.StripPrefix("/public", public))
	r.Handle("/applications/", http.StripPrefix("/applications", authMiddleware(applications, tokenMaker)))
	r.Handle("/blogs/", http.StripPrefix("/blogs", authMiddleware(blogs, tokenMaker)))
	r.Handle("/educations/", http.StripPrefix("/educations", authMiddleware(educations, tokenMaker)))
	r.Handle("/experiences/", http.StripPrefix("/experiences", authMiddleware(experiences, tokenMaker)))
	r.Handle("/links/", http.StripPrefix("/links", authMiddleware(links, tokenMaker)))
	r.Handle("/files/", http.StripPrefix("/files", authMiddleware(files, tokenMaker)))
	r.Handle("/logs/", http.StripPrefix("/logs", authMiddleware(logs, tokenMaker)))
	r.Handle("/messages/", http.StripPrefix("/messages", authMiddleware(messages, tokenMaker)))
	r.Handle("/users/", http.StripPrefix("/users", authMiddleware(users, tokenMaker)))
	r.Handle("/portfolios/", http.StripPrefix("/portfolios", authMiddleware(portfolios, tokenMaker)))
	r.Handle("/services/", http.StripPrefix("/services", authMiddleware(services, tokenMaker)))
	r.Handle("/testimonials/", http.StripPrefix("/testimonials", authMiddleware(testimonials, tokenMaker)))
	r.Handle("/tokens/", http.StripPrefix("/tokens", tokens))
	r.Handle("/website/", http.StripPrefix("/website", website))
}

func stack() models.MiddleWare {
	return createStack(
		logging,
		corsMiddleware,
		origPath,
	)
}

func Start(addr string) error {
	log.Printf("Server started and listening on port :%s", addr)
	return http.ListenAndServe(fmt.Sprintf(":%s", addr), stack()(r))
}
