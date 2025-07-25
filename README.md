# 📌 Dynamic Todo API with PostgreSQL (GET endpoint)

This is a simple, beginner-friendly Go backend project that connects to a PostgreSQL database and supports **dynamic filtering** using query parameters like `title`, `done`, and `limit`.

---

## 📦 Features

- 🔍 Search todos by title (`ILIKE`)
- ✅ Filter by completion status
- 📄 Limit number of results
- 📦 Clean and simple project structure
- 🛠️ Easy to understand for beginners

---

## 🛠️ Technologies Used

- Go (Golang)
- PostgreSQL
- `net/http`
- `database/sql`
- `github.com/lib/pq`

---

## 🚀 How to Run the Project

### 1. Clone the repo

```bash
git clone https://github.com/ssdtoshkentov/DynamicGet
cd DynamicGet
```

### 2. Set up PostgreSQL

Create a database called `todo_db` and a table named `todos`:

```sql
CREATE DATABASE todo_db;

\c todo_db

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT,
    done BOOLEAN
);

-- Sample data
INSERT INTO todos (title, done) VALUES
('Learn Go', false),
('Build an API', true),
('Read a book', false);
```

### 3. Update connection string

In `main.go`, update your PostgreSQL connection string:

```go
connStr := "user=Abduvali password=ssdAbduvali dbname=todo_db sslmode=disable"
```

Replace `Abduvali` and `ssdAbduvali` with your actual PostgreSQL credentials.

### 4. Run the API

```bash
go run main.go
```

API will be running on: `http://localhost:8080/todos`

---

## 📝 Query Parameters

| Parameter | Type    | Description                            |
|-----------|---------|----------------------------------------|
| `title`   | string  | Partial match (case-insensitive)       |
| `done`    | boolean | Filter by completion status (`true`)   |
| `limit`   | int     | Limit the number of results returned   |

---

## 🔍 Example Requests

```http
GET /todos?title=go
GET /todos?done=true
GET /todos?limit=5
GET /todos?title=book&done=false&limit=10
```

---

## 📂 File Structure

```
your-repo-name/
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

## 🙌 Author

Made with ❤️ by [Abduvali]

---

## 🧠 License

This project is licensed under the MIT License.
