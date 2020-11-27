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

func AddItem(i *Item) {
	i.ID = getNextId()
	itemList = append(itemList, i)
}

func getNextId() int {
	li := itemList[len(itemList)-1]
	return li.ID + 1
}

type Items []*Item

func (i *Item) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

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
