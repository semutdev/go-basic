package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	jumlahKlik := 0

	http.HandleFunc("/klik", func(w http.ResponseWriter, r *http.Request) {
		jumlahKlik++

		htmlBalasan := fmt.Sprintf(`<button hx-post="/klik" hx-swap="outerHTML">Wah, sudah diklik %d kali!</button>`, jumlahKlik)

		fmt.Fprint(w, htmlBalasan)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
