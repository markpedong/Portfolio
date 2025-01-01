# Portfolio Backend Service  

A robust backend web service built with **Go**, leveraging **Prisma** for database migrations, **sqlx** for database communication, and **gRPC** for high-performance service communication. The project is containerized with **Docker** for seamless deployment and scalability.  

## Key Features  
- **Customizable Content Management System (CMS):**  
  - Manage and toggle visibility for portfolio sections such as:  
    - Applications Stack  
    - Blogs  
    - Education  
    - Experience  
    - Portfolio Items  
    - Testimonials  
  - Tailor content dynamically to meet user requirements.  

- **Messaging Functionality:**  
  - Includes a message-sending feature for user communication.  

- **Containerized Deployment:**  
  - Services are containerized using Docker for consistent and portable deployment.  

## Technologies Used  
- **Programming Language:** Go  
- **Database Management:** Prisma, sqlx  
- **Communication Protocol:** gRPC  
- **Containerization:** Docker  

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

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Build and Run the Services**
   ```bash
   docker-compose up -d
   ```
4. **Run Prisma Migrations**
   ```bash   
   prisma migrate dev --name init
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