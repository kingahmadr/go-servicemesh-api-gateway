package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Definisikan struct untuk Review
type Review struct {
	ProductID string `json:"productId"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}

// Buat database review in-memory
var reviews = map[string][]Review{
	"1": {
		{ProductID: "1", Rating: 5, Comment: "Sangat membantu untuk pemula!"},
		{ProductID: "1", Rating: 4, Comment: "Penjelasannya jelas."},
	},
	"2": {
		{ProductID: "2", Rating: 5, Comment: "Enak dipakai mengetik."},
	},
}

func reviewsHandler(w http.ResponseWriter, r *http.Request) {
	productID := strings.TrimPrefix(r.URL.Path, "/reviews/")
	if productID == "" {
		http.Error(w, "Product ID harus disertakan", http.StatusBadRequest)
		return
	}

	if productReviews, ok := reviews[productID]; ok {
		json.NewEncoder(w).Encode(productReviews)
	} else {
		// Kembalikan array kosong jika tidak ada review
		json.NewEncoder(w).Encode([]Review{})
	}
}

func main() {
	http.HandleFunc("/reviews/", reviewsHandler)

	log.Println("Review service berjalan di port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
