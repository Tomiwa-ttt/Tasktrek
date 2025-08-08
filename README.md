# Tasktrek 🗂️

**Tasktrek** is a simple task management API built with Go and MongoDB. It allows users to create, read, update, and delete tasks — perfect for learning backend development, RESTful API design, and deployment with Render.

---

## 🚀 Features

- Create new tasks
- Retrieve all or a single task
- Update existing tasks
- Delete tasks
- Connects to MongoDB Atlas
- Deployed on [Render](https://render.com)

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Database**: MongoDB (via MongoDB Atlas)
- **Framework**: Native Go net/http
- **Environment**: Render (deployment)
- **Package Management**: Go modules

---

## 📁 Folder Structure



Tasktrek/
├── config/
│   └── db.go          # MongoDB connection setup
├── controllers/
│   └── task.go        # Handler functions
├── models/
│   └── task.go        # Task model
├── routes/
│   └── routes.go      # API route definitions
├── main.go            # Entry point
├── go.mod             # Dependencies and Go version
├── go.sum
└── .env               # MongoDB URI (keep secret!)

``

---

## ⚙️ Setup & Run Locally

### 1. Clone the repo

```bash
git clone https://github.com/Tomiwa-ttt/Tasktrek.git
cd Tasktrek
````

### 2. Create a `.env` file

```env
MONGODB_URI=your-mongo-uri-here
```

> Replace with your actual MongoDB connection string.

### 3. Run the server

```bash
go mod tidy
go run main.go
```

Server should start on `http://localhost:8080`.

---

## 🧪 API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/tasks`      | Get all tasks     |
| GET    | `/tasks/{id}` | Get a task by ID  |
| POST   | `/tasks`      | Create a new task |
| PUT    | `/tasks/{id}` | Update a task     |
| DELETE | `/tasks/{id}` | Delete a task     |

---

## 🌐 Deployment

Deployed using [Render](https://render.com). The server automatically builds and redeploys on new commits.

---

## 📜 License

MIT License — free to use, modify, and distribute.

---

## ✨ Author

Built by [@Tomiwa-ttt](https://github.com/Tomiwa-ttt)

```

---

Let me know if you want a version with screenshots, usage examples (like sample JSON), or documentation badges.
```
