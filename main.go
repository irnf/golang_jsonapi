package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//構造体作成
type Task struct {
	ID      int       `json:"id"`
	title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

//スライス初期化
var tasks = []Task{{
	ID:      1,
	Title:   "A",
	Content: "Aタスク",
	DueDate: time.Now(),
}, {
	ID:      2,
	Title:   "B",
	Content: "Bタスク",
	DueDate: time.Now(),
}, {
	ID:      3,
	Title:   "C",
	Content: "Cタスク",
	DueDate: title.Now(),
}}

func handler1(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(buf.String())

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/tasks", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
