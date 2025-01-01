package models

import (
	"net/http"
	"time"

	"github.com/lib/pq"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type MiddleWare func(http.Handler) http.Handler

type Messages struct {
	ID        string     `db:"id" json:"id"`
	Name      string     `db:"name" json:"name" validate:"required"`
	Email     string     `db:"email" json:"email" validate:"required"`
	Message   string     `db:"message" json:"message" validate:"required"`
	Status    int        `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Portfolios struct {
	ID        string         `db:"id" json:"id"`
	Title     string         `db:"title" json:"title" validate:"required"`
	Tech      pq.StringArray `db:"tech" json:"tech" validate:"required"`
	Link      string         `db:"link" json:"link" validate:"required"`
	Image     string         `db:"image" json:"image" validate:"required"`
	Status    int            `db:"status" json:"status"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time     `db:"deleted_at" json:"deleted_at"`
}

type Blogs struct {
	ID          string     `db:"id" json:"id"`
	Title       string     `db:"title" json:"title" validate:"required"`
	Date        string     `db:"date" json:"date" validate:"required"`
	Description string     `db:"description" json:"description" validate:"required"`
	Link        string     `db:"link" json:"link" validate:"required"`
	Image       string     `db:"image" json:"image" validate:"required"`
	Status      int        `db:"status" json:"status"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Testimonials struct {
	ID          string     `json:"id" db:"id"`
	Author      string     `json:"author" validate:"required" db:"author"`
	Description string     `json:"description" validate:"required" db:"description"`
	Image       string     `json:"image" validate:"required" db:"image"`
	Job         string     `json:"job" validate:"required" db:"job"`
	Status      int        `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type Color struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Theme     string `json:"theme"`
	Title     string `json:"title"`
	Color     string `json:"color"`
	WebsiteID string `json:"website_id"`
}

type Website struct {
	ID        string     `db:"id" json:"id"`
	Status    int        `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Application struct {
	ID        string     `db:"id" json:"id"`
	Image     string     `db:"image" json:"image" validate:"required"`
	Name      string     `db:"name" json:"name" validate:"required"`
	Status    int        `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Files struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	File      string    `db:"file" json:"file"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type EduSkill struct {
	ID          string `db:"id" json:"id"`
	EducationID string `db:"education_id" json:"education_id"`
	Name        string `db:"name" json:"name"`
	Percentage  int    `db:"percentage" json:"percentage"`
}

type Education struct {
	ID          string     `db:"id" json:"id"`
	School      string     `db:"school" json:"school"`
	Course      string     `db:"course" json:"course"`
	Started     string     `db:"started" json:"started"`
	Ended       string     `db:"ended" json:"ended"`
	Description string     `db:"description" json:"description"`
	Status      int        `db:"status" json:"status"`
	Skills      []EduSkill `json:"skills"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}

type ExpSkill struct {
	ID           string `db:"id" json:"id"`
	ExperienceID string `db:"experience_id" json:"experience_id"`
	Name         string `db:"name" json:"name"`
	Percentage   int    `db:"percentage" json:"percentage"`
}

type Experiences struct {
	ID           string         `json:"id" db:"id"`
	Company      string         `json:"company" db:"company"  validate:"required"`
	Title        string         `json:"title" db:"title" validate:"required"`
	Location     string         `json:"location" db:"location" validate:"required"`
	Started      string         `json:"started" db:"started" validate:"required"`
	Ended        string         `json:"ended" db:"ended" validate:"required"`
	Skills       []ExpSkill     `json:"skills" db:"skills"`
	Descriptions pq.StringArray `json:"descriptions" db:"descriptions"`
	Status       int            `json:"status" db:"status"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time     `json:"deleted_at" db:"deleted_at"`
}

type Links struct {
	ID        string     `db:"id" json:"id"`
	Link      string     `db:"link" json:"link" validate:"required"`
	Type      string     `db:"type" json:"type" validate:"required"`
	Status    int        `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Services struct {
	ID          string     `db:"id" json:"id"`
	Title       string     `db:"title" json:"title" validate:"required"`
	Description string     `db:"description" json:"description" validate:"required"`
	Logo        string     `db:"logo" json:"logo" validate:"required"`
	Status      int        `db:"status" json:"status"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Users struct {
	ID             string     `db:"id" json:"id"`
	FirstName      string     `db:"first_name" json:"first_name"`
	LastName       string     `db:"last_name" json:"last_name"`
	Phone          string     `db:"phone" json:"phone"`
	Address        string     `db:"address" json:"address"`
	Description    string     `db:"description" json:"description"`
	Email          string     `db:"email" json:"email"`
	Username       string     `db:"username" json:"username"`
	Password       string     `db:"password" json:"password"`
	ResumePDF      string     `db:"resume_pdf" json:"resume_pdf"`
	ResumeDocx     string     `db:"resume_docx" json:"resume_docx"`
	IsDownloadable int        `db:"isdownloadable" json:"isdownloadable"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at"`
}

type Session struct {
	ID           string     `json:"id" db:"id"`
	UserID       string     `json:"user_id" db:"user_id"`
	Email        string     `json:"email" db:"email"`
	RefreshToken string     `json:"refresh_token" db:"refresh_token"`
	IsRevoked    bool       `json:"is_revoked" db:"is_revoked"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	ExpiresAt    *time.Time `json:"expires_at" db:"expires_at"`
}
