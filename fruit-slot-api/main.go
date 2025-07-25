package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
)

var fruits = []string{"Cherry", "Lemon", "Orange", "Grape", "Watermelon", "Pineapple"}

// getRandomFruits generates 3 random fruits securely
func getRandomFruits() ([]string, error) {
	result := make([]string, 3)
	for i := 0; i < 3; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(fruits))))
		if err != nil {
			return nil, err
		}
		result[i] = fruits[num.Int64()]
	}
	return result, nil
}

// checkWin determines if at least 2 fruits are the same
func checkWin(spin []string) bool {
	counts := make(map[string]int)
	for _, fruit := range spin {
		counts[fruit]++
		if counts[fruit] >= 2 {
			return true
		}
	}
	return false
}

// handler for /play
func playHandler(w http.ResponseWriter, r *http.Request) {
	spin, err := getRandomFruits()
	if err != nil {
		http.Error(w, "Failed to generate fruits", http.StatusInternalServerError)
		return
	}

	message := "Try again!"
	if checkWin(spin) {
		message = "You win!"
	}

	response := map[string]interface{}{
		"fruits":  spin,
		"message": message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/play", playHandler)
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
