package data

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
}

func (i *Item) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(i)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
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
		SKU:         "abc-def-ghi",
		CreatedOn:   time.Now().UTC().String(),
	},
	&Item{
		ID:          2,
		Name:        "Bed",
		Description: "Simple Bed",
		Price:       15000.00,
		SKU:         "agg-bdd-ddd",
		CreatedOn:   time.Now().UTC().String(),
	},
}
