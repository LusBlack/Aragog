package main

import (
	"fmt"
	"net/http"

	"github.com/LusBlack/Aragog/crawler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{vin}", crawler.CrawlerForEngine)

	fmt.Println("We're Oscar Mike. localhost:8000/{vin}")
	http.ListenAndServe(":8000", router)
}
