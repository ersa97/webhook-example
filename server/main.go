package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var callbackURL = "http://localhost:8080/webhook"

func main() {
	http.HandleFunc("/callback", callback)
	fmt.Println("Server Started on 8000")
	http.ListenAndServe(":8000", nil)
}

func callback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	payload := []byte(`("Code":200, "Message":"Hi From Server")`)

	//send the callback with webhook
	resp, err := http.Post(callbackURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending webhook:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Webhook response:", resp.Status)

	w.WriteHeader(http.StatusOK)
}
