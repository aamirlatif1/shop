package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aamirlatif1/shop/handlers"
)

func main() {
	l := log.New(os.Stdout, "rgs ", log.LstdFlags)

	sm := http.NewServeMux()

	sm.Handle("/", handlers.NewHello(l))
	sm.Handle("/items", handlers.NewItem(l))

	http.ListenAndServe(":8080", sm)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
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
