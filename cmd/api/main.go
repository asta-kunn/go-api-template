package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Membuat route untuk path '/'
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	// Mengambil port dari environment variable (Sangat penting untuk proses deploy)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port jika dijalankan secara lokal
	}

	fmt.Printf("Server berjalan di port :%s\n", port)

	// Menjalankan server menggunakan port yang sudah disesuaikan
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}