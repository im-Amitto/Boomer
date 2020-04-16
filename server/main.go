package main

import (
	"fmt"
	"log"
	"net/http"

	"./router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 3200...")
	log.Fatal(http.ListenAndServe(":3200", r))
}
