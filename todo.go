package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type Task struct {
	ID int
	Title string
	Done bool
}

var tasks []Task

var nextID int = 1

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

// main function
func main() {
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
			_, err := fmt.Sscanf(strings.TrimSpace(idStr), "%d", &id)

			if err != nil {
				fmt.Println("error: Input harus berupa angka.")
			}else{
				completeTask(id)
			}
		case "4":
			fmt.Println("Sampai jumpa!")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}

	}
}