package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

func main() {
	http.HandleFunc("/todos", getTodosHandler)

		fmt.Println("Server running on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodosHandler(w http.ResponseWriter, r *http.Request)
//1. sorov parametrlarni olish
titleParam := r.URL.Query().Get("title")
doneParam := r.URL.Query().Get("done")
limitParam := r.URL.Query().Get("limit")


// 2. PostgreSQLga ulanish
	connStr := "user=Abduvali password=ssdToshkentov dbname=todo_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 3.Dynamic SQL query yasash
	query := `Select id, title, done FROM todos WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	if titleParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			http.Error(w, "Invalid Limit", http.StatusBadRequest)
			return
		}
		query += fmt.Sprintf(" Limit %d", limit)
	}

	// 4. So‘rov yuborish
		rows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, "DB query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()


		// 5. Natijalarni o‘qish
		todos := []Todo{}
		for rows.Next() {
			var t Todo
			err := rows.Scan(&t.ID, &t.Title, &t.Done)
			if err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				return
			}
			todos = append(todos, t)
		}


		// 6. JSON qilib chiqarish
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
}
