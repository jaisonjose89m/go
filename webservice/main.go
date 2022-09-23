package main

import (
	"fmt"
	"github.com/jaisonjose89m/go/webservice/controllers"
	"net/http"
)

func main() {
	fmt.Println("Starting the server")
	controllers.RegisterController()
	http.ListenAndServe(":8082", nil)
}
