package models

import (
	"time"

	"github.com/lib/pq"
)

type RenewAccessTokenRes struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type PApplications struct {
	ID    string `db:"id" json:"id"`
	Image string `db:"image" json:"image" validate:"required"`
	Name  string `db:"name" json:"name" validate:"required"`
}

type PBlogs struct {
	ID          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title" validate:"required"`
	Date        string `db:"date" json:"date" validate:"required"`
	Description string `db:"description" json:"description" validate:"required"`
	Link        string `db:"link" json:"link" validate:"required"`
	Image       string `db:"image" json:"image" validate:"required"`
}

type PLinks struct {
	ID   string `db:"id" json:"id"`
	Link string `db:"link" json:"link" validate:"required"`
	Type string `db:"type" json:"type" validate:"required"`
}

type PServices struct {
	ID          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title" validate:"required"`
	Description string `db:"description" json:"description" validate:"required"`
	Logo        string `db:"logo" json:"logo" validate:"required"`
}

type PTestimonials struct {
	ID          string `json:"id" db:"id"`
	Author      string `json:"author" validate:"required" db:"author"`
	Description string `json:"description" validate:"required" db:"description"`
	Image       string `json:"image" validate:"required" db:"image"`
	Job         string `json:"job" validate:"required" db:"job"`
}

type PPortfolios struct {
	ID    string         `db:"id" json:"id"`
	Title string         `db:"title" json:"title" validate:"required"`
	Tech  pq.StringArray `db:"tech" json:"tech" validate:"required"`
	Link  string         `db:"link" json:"link" validate:"required"`
	Image string         `db:"image" json:"image" validate:"required"`
}

type PEduSkill struct {
	ID         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Percentage int    `db:"percentage" json:"percentage"`
}

type PEducation struct {
	ID          string      `db:"id" json:"id"`
	School      string      `db:"school" json:"school"`
	Course      string      `db:"course" json:"course"`
	Started     string      `db:"started" json:"started"`
	Ended       string      `db:"ended" json:"ended"`
	Description string      `db:"description" json:"description"`
	Skills      []PEduSkill `json:"skills"`
}

type PExpSkill struct {
	ID         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Percentage int    `db:"percentage" json:"percentage"`
}

type PExperiences struct {
	ID           string         `json:"id" db:"id"`
	Company      string         `json:"company" db:"company"  validate:"required"`
	Title        string         `json:"title" db:"title" validate:"required"`
	Location     string         `json:"location" db:"location" validate:"required"`
	Started      string         `json:"started" db:"started" validate:"required"`
	Ended        string         `json:"ended" db:"ended" validate:"required"`
	Skills       []PExpSkill    `json:"skills" db:"skills"`
	Descriptions pq.StringArray `json:"descriptions" db:"descriptions"`
}

type PublicUser struct {
	ID             string `db:"id" json:"id"`
	FirstName      string `db:"first_name" json:"first_name"`
	LastName       string `db:"last_name" json:"last_name"`
	Phone          string `db:"phone" json:"phone"`
	Address        string `db:"address" json:"address"`
	Description    string `db:"description" json:"description"`
	Email          string `db:"email" json:"email"`
	Username       string `db:"username" json:"username"`
	ResumePDF      string `db:"resume_pdf" json:"resume_pdf"`
	ResumeDocx     string `db:"resume_docx" json:"resume_docx"`
	IsDownloadable int    `db:"isdownloadable" json:"isdownloadable"`
}
