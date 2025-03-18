# Codwikz
A backend system for an online coding platform where users can solve programming problems, submit solutions, and track their progress.

## 🚀 Features

- User Authentication – Secure login & registration using JWT & OAuth.
- Problem Management – APIs to create, update, and fetch coding problems.
- Code Submission – Handles and stores user submissions efficiently.
- Real-Time Leaderboard – Uses WebSockets for live ranking updates.
- Code Execution Engine – Docker-based execution in an isolated environment.

## 🛠️ Tech Stack
- Backend: Golang (Gin/Fiber)
- Database: PostgreSQL
- API Communication: REST & gRPC
- Containerization: Docker

## ⚙️ Setup & Installation
- Clone the Repository
  
git clone https://github.com/rajritwika1/codwikz.git

cd codwikz

## Setup PostgreSQL Database

CREATE DATABASE codwikz_db;

## Configure Environment Variables (.env file)

PORT=8080  
DB_URL=postgres://username:password@localhost:5432/codwikz_db?sslmode=disable  
JWT_SECRET=your_secret_key  

## Run the Backend Server

go run main.go

## 📌 API Endpoints
- User Authentication
  
POST /register – Register a new user

POST /login – Login and get JWT token

- Problems
  
GET /problems – Fetch all coding problems

POST /problems – Add a new problem

- Code Submission
  
POST /submit – Submit a solution for evaluation

## 📌 To-Do List
- [ ] Implement test cases
- [ ] Add admin panel for problem management
- [ ] Improve error handling
 
## 🛠️ Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

