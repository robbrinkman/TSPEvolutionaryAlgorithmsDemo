package main
import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../frontend/app"))
	http.Handle("/", fs)

	log.Println("Listening...")

	// TODO add router

	// TODO add handler
	http.ListenAndServe(":3000", nil)
}
