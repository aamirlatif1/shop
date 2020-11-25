package handlers

import (
	"log"
	"net/http"

	"github.com/aamirlatif1/shop/data"
)

type Items struct {
	l *log.Logger
}

func NewItem(l *log.Logger) *Items {
	return &Items{l}
}

func (i *Items) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		i.GetItems(rw, r)
		return
	}

	rw.WriteHeader(http.StatusNotImplemented)
}

func (i *Items) GetItems(rw http.ResponseWriter, r *http.Request) {
	itemList := data.GetItems()
	err := itemList.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marschal data", http.StatusInternalServerError)
		return
	}
}