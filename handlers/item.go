package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aamirlatif1/shop/data"
	"github.com/gorilla/mux"
)

type KeyItem struct{}

type Items struct {
	l *log.Logger
}

func NewItem(l *log.Logger) *Items {
	return &Items{l}
}

func (i *Items) GetItems(rw http.ResponseWriter, r *http.Request) {
	itemList := data.GetItems()
	err := itemList.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marschal data", http.StatusInternalServerError)
		return
	}
}

func (i *Items) AddItem(rw http.ResponseWriter, r *http.Request) {
	it := r.Context().Value(KeyItem{}).(data.Item)
	i.l.Printf("%v", it)
	data.AddItem(&it)
}

func (i *Items) UpdateItem(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	it := r.Context().Value(KeyItem{}).(data.Item)
	i.l.Printf("%v", it)

	i.l.Printf("%v", vars)
}

func (i *Items) MiddlewareValidateItem(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		it := data.Item{}
		err := it.FromJson(r.Body)
		if err != nil {
			i.l.Println("[ERROR] deserializing item", err)
			http.Error(rw, "Unable to read item", http.StatusBadRequest)
			return
		}

		err = it.Validate()
		if err != nil {
			i.l.Println("[ERROR] Required value missing ", err)
			http.Error(rw, fmt.Sprintf("Validation Failed : %v", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyItem{}, it)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
