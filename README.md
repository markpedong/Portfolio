# Portfolio Backend Service and Admin Interface

A robust full-stack web service with a **Go** backend, **React** admin interface, and containerized with **Docker** for seamless deployment and scalability. This project includes a **CMS** for managing portfolio sections, message-sending functionality, and a backend API that communicates with a database using **Prisma** and **sqlx**.

The **frontend** is built with **React** using **Vite** for fast development, and the app is served via **Nginx** with a hot-reload feature during development.

## Key Features

- **Backend (Go):**

  - CMS for managing portfolio sections (Applications Stack, Blogs, Education, Experience, Portfolio Items, Testimonials).
  - Includes a messaging feature for communication with users.
  - gRPC for efficient service communication.
  - Database management using **Prisma** and **sqlx**.

- **Frontend (React):**

  - Admin Interface built with **React** and **Vite**.
  - Hot-reload enabled via **Nginx** for smooth development experience.
  - Manage portfolio items and other sections directly from the admin interface.

- **Containerized Deployment:**
  - Backend and frontend are containerized using **Docker** and **Docker Compose** for consistent and easy deployment.

## Technologies Used

- **Backend (Go):** Go, gRPC, Prisma, sqlx
- **Frontend (React):** React, Vite, Nginx
- **Containerization:** Docker, Docker Compose

### Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/)
- [Node.js](https://nodejs.org/) (for Prisma CLI)

### Steps

1. **Clone the Repository**  
   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   add necessary environment variables
   ```

2. **Install Dependencies**

   ```bash
   cd server && go mod tidy
   cd client && npm install
   ```

3. **Run Prisma Migrations**

   ```bash
   cd server && npx prisma migrate dev --name init
   ```

4. **Build and Run the Services**

   ```bash
    cd client && docker-compose up -d
    cd server && docker-compose up -d
   ```

5. **Access the Admin Interface**
   Navigate to http://localhost:6602 in your web browser to access the admin interface.

6. **Stop and Remove the Services**
   ```bash
   cd client && docker-compose down
   cd server && docker-compose down
   ```

## Troubleshooting

- **Prisma Migration Error:**

  - Check if the project has been initialized with the correct Prisma schema.
  - If not, run the `npx prisma migrate dev --name init` command to initialize the database.

- **Docker Compose Error:**

  - Ensure you have Docker and Docker Compose installed on your system.

- **Database Connection Error:**

  - Check if the database is running and accessible from your machine.

- **Containerization Error:**
  - If you're unable to run the Docker containers, ensure that you have Docker installed and are following the instructions provided in the "Getting Started" section.

## License

This project is licensed under the [MIT License](LICENSE).
