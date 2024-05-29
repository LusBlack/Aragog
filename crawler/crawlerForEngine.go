package crawler

import (
	"fmt"
	"net/http"

	//"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func CrawlerForEngine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vin := vars["vin"]

	//create new colly collector
	// c := colly.NewCollector(
	// 	colly.AllowURLRevisit(),
	// 	colly.MaxDepth(100),
	// )

	fmt.Fprintf(w, "You used this vin: %s", vin)

}
