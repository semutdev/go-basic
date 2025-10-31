package main

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
)

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

var tasks []Task
var nextID int = 1
const fileName = "tasks.json"

func saveTasksToFile() error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("gagal encode JSON: %w", err)
	}

	err = os.WriteFile(fileName, data, 0664)
	if err != nil {
		return fmt.Errorf("gagal menyimpan file: %w", err)
	}

	// simpan next id dengan cara ambil paling tinggi id nya
	if len(tasks) > 0 {
		maxID := 0
		for _, t := range tasks {
			if t.ID > maxID {
				maxID = t.ID
			}
			nextID = maxID + 1
		}
	} else {
		nextID = 1
	}

	return nil
}

func loadTasksFromFile() error {
	data, err := os.ReadFile(fileName)

	if os.IsNotExist(err) {
		fmt.Println("File tasks.json tidak ditemukan")
		tasks = []Task{}
		return nil
	}

	if err != nil {
		return fmt.Errorf("gagal membaca file: %w", err)
	}

	// json ke slice
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return fmt.Errorf("Gagal decode JSON: %w", err)
	}

	if len(tasks) > 0 {
		maxID := 0
		for _, t := range tasks {
			if t.ID > maxID {
				maxID = t.ID
			}
		}
		nextID = maxID + 1
	}

	fmt.Println("Berhasil memuat", len(tasks), "tugas dari", fileName)
	return nil
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(tasks)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case "POST":
			var newTask Task

			err := json.NewDecoder(r.Body).Decode(&newTask)
			if err != nil {
				http.Error(w, "Bad request json not valid", http.StatusBadRequest)
				return
			}

			// validasi title tidak boleh kosong
			if newTask.Title == "" {
				http.Error(w, "title tidak boleh kosong", http.StatusBadRequest)
				return
			}

			newTask.ID = nextID
			nextID++
			newTask.Done = false

			tasks = append(tasks, newTask)

			if err := saveTasksToFile(); err != nil {
				http.Error(w, "Internal Server error: gagal menyimpan data", http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(newTask)
		default:
			http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

// main function
func main() {
	// init file
	if err := loadTasksFromFile(); err != nil {
		fmt.Println("Error saat memuat tugas:", err)
	}

	http.HandleFunc("/tasks", taskHandler)

	port := ":8080"
	fmt.Println("server run on port" + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Gagal menyalakan server:", err)
	}
}