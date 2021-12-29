package functions

import "fmt"

func InvokeFunc() {
	_, err := startWebServer(5000, 3)
	fmt.Println(err)
}
func startWebServer(port, retryCount int) (int, error) {
	fmt.Println("Starting server")
	fmt.Println("Started server")
	return port, nil
}
