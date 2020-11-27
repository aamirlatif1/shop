package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aamirlatif1/shop/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "rgs ", log.LstdFlags)

	r := mux.NewRouter()

	it := handlers.NewItem(l)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/items", it.GetItems)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/items", it.AddItem)
	postRouter.Use(it.MiddlewareValidateItem)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/items/{id:[0-9]+}", it.UpdateItem)
	putRouter.Use(it.MiddlewareValidateItem)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":8080", r)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	l.Println("Received terminate, shutting down gracefully", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
