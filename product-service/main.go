package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Definisikan struct untuk Product
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Buat database in-memory sederhana
var products = []Product{
	{ID: "1", Name: "Buku Pemrograman Go", Price: 150000},
	{ID: "2", Name: "Keyboard Mekanikal", Price: 850000},
	{ID: "3", Name: "Headphone", Price: 500000},
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	// Cek apakah request untuk detail produk atau semua produk
	id := strings.TrimPrefix(r.URL.Path, "/products/")
	if id != "" && id != "products/" {
		// Cari produk berdasarkan ID
		for _, p := range products {
			if p.ID == id {
				json.NewEncoder(w).Encode(p)
				return
			}
		}
		http.NotFound(w, r)
		return
	}

	// Kembalikan semua produk
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productsHandler) // Menangani path dengan trailing slash

	log.Println("Product service berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
