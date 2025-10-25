package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"encoding/json"
	"strconv"
)

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

var tasks []Task
var nextID int = 1
const fileName = "tasks.json"

func addTask(title string) {
	task := Task{
		ID: nextID,
		Title: title,
		Done: false,
	}

	tasks = append(tasks, task)

	nextID++
	fmt.Println("Tugas berhasil ditambahkan:", title)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas.")
		return
	}

	fmt.Println("Daftar Tugas:")
	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "X"
		}

		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
	}
}

func completeTask(id int) {
	found := false
	for i := range tasks{
		if tasks[i].ID == id {
			tasks[i].Done = true
			found = true
			fmt.Println("Tugas ditandai selesai:", tasks[i].Title)
			break
		}
	}

	if !found {
		fmt.Println("Error: Tugas dengan ID", id, "tidak ditemukan")
	}
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

// main function
func main() {
	// init file
	if err := loadTasksFromFile(); err != nil {
		fmt.Println("Error saat memuat tugas:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selamat datang di app todo list")

	for{
		// List menu
		fmt.Println("\n---- Menu ----")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat semua tugas")
		fmt.Println("3. Tandai selesai")
		fmt.Println("4. Keluar")
		fmt.Println("Pilih opsi: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("Masukan nama tugas:")
			title, _ := reader.ReadString('\n')

			addTask(strings.TrimSpace(title))
		case "2":
			listTasks()
		case "3":
			fmt.Println("Masukan ID tugas yang selesai:")
			idStr, _ := reader.ReadString('\n')
			var id int
			// string ke int
			id, err := strconv.Atoi(strings.TrimSpace(idStr))

			if err != nil {
				fmt.Println("error: Input harus berupa angka.")
			}else{
				completeTask(id)
			}
		case "4":
			if err := saveTasksToFile(); err != nil {
				fmt.Println("Error saat menyimpan tugas:", err)
			}else{
				fmt.Println("Tugas berhasil disimpan. sampai jumpa!")
			}
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}

	}
}