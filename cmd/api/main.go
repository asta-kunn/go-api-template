package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)

	// Menjalankan server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
