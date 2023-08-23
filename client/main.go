package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	fmt.Println("Server Client started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println("Received webhook:", string(body))
	//do something with the callbacks
	w.WriteHeader(http.StatusOK)
}
