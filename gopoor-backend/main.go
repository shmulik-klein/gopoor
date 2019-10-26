package main

import (
	"log"
	"net/http"

	"github.com/unrolled/render"
)

func main() {
	http.HandleFunc("/", PurchaseHandler(&render.Render{}))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
