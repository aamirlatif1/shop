package data

import (
	"encoding/json"
	"io"
	"time"
)

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"-"`
}

func GetItems() Items {
	return itemList
}

type Items []*Item

func (i *Items) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

var itemList = []*Item{
	&Item{
		ID:          1,
		Name:        "Chair",
		Description: "Arm Chair",
		Price:       5000.00,
		CreatedOn:   time.Now().UTC().String(),
	},
	&Item{
		ID:          2,
		Name:        "Bed",
		Description: "Simple Bed",
		Price:       15000.00,
		CreatedOn:   time.Now().UTC().String(),
	},
}
