package utils

const (
	UpdateLink        = "UPDATE links SET link=:link"
	UpdateService     = "UPDATE services SET title=:title, description=:description, logo=:logo"
	UpdatePortfolio   = "UPDATE portfolios SET title=:title, tech=:tech, link=:link, image=:image"
	UpdateApplication = "UPDATE applications SET name = :name, image = :image"
	UpdateUser        = "UPDATE users SET first_name=:first_name, last_name=:last_name, phone=:phone, address=:address, description=:description, email=:email, username=:username, resume_pdf=:resume_pdf, resume_docx=:resume_docx, isdownloadable=:isdownloadable"
	UpdateWebsite     = "UPDATE website SET status=:status"
	UpdateTestimonial = "UPDATE testimonials SET author=:author, description=:description, image=:image, job=:job"
	UpdateBlog        = "UPDATE blogs SET title=:title, date=:date, description=:description, link=:link, image=:image"

	CreateService   = "INSERT INTO services (title, description, logo) VALUES (:title, :description, :logo)"
	CreateLink      = "INSERT INTO links (type, link) VALUES (:type, :link)"
	CreateApp       = "INSERT INTO applications (name, image) VALUES (:name, :image)"
	CreatePortfolio = "INSERT INTO portfolios (title, tech, link, image) VALUES (:title, :tech, :link, :image)"
	CreateUser      = `
		INSERT INTO users (first_name, last_name, phone, address, description, email, username, password, resume_pdf, resume_docx, isdownloadable)
		VALUES (:first_name, :last_name, :phone, :address, :description, :email, :username, :password, :resume_pdf, :resume_docx, :isdownloadable)
	`
	CreateTestimonial = "INSERT INTO testimonials (author, description, image, job) VALUES (:author, :description, :image, :job)"
	CreateBlog        = `INSERT INTO blogs (title, date, description, link, image) VALUES (:title, :date, :description, :link, :image)`
	CreateMessage     = `INSERT INTO messages (name, email, message) VALUES (:name, :email, :message)`
)
