package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aamirlatif1/shop/handlers"
)

func main() {
	l := log.New(os.Stdout, "rgs", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	http.ListenAndServe(":8080", sm)
}
