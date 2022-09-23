package controllers

import (
	"encoding/json"
	"net/http"
)

func RegisterController() {
	controller := newUserController()
	http.Handle("/users", *controller)
	http.Handle("/users/", *controller)
	http.Handle("/api/v1/credentials/query/keystores", *controller)
}

func EncodeAsJsonString(data interface{}, w http.ResponseWriter) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
