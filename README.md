# Portfolio Backend Service

A robust backend web service built with **Go**, leveraging **Prisma** for database migrations, **sqlx** for database communication, and **gRPC** for high-performance service communication. This project is containerized using **Docker** for seamless deployment and scalability.

The backend includes a **Content Management System (CMS)** to manage portfolio sections dynamically, such as Applications Stack, Blogs, Education, Experience, Portfolio Items, and Testimonials. Additionally, it features a message-sending function for user communication.

## Key Features

- **Backend:**

  - CMS for managing portfolio sections (Applications Stack, Blogs, Education, Experience, Portfolio Items, Testimonials).
  - Messaging feature for communication.
  - gRPC for efficient and scalable service-to-service communication.
  - Database migrations and management using **Prisma** and **sqlx**.

- **Containerized Deployment:**
  - Fully containerized using Docker and Docker Compose for consistency and ease of deployment.

## Technologies Used

- **Programming Language:** Go
- **Database Management:** Prisma, sqlx
- **Communication Protocol:** gRPC
- **Containerization:** Docker

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/)

### Installation

1. **Clone the Repository**  
   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   ```

2. **Install Dependencies and Build**  
   Navigate to the project directory and install dependencies:

   ```bash
   go mod tidy
   prisma migrate dev --schema=./db/schema.prisma
   ./build-server.sh
   ```