package data

import (
	"time"

	"github.com/permalik/temp_rest_go/internal/validator"
)

type Item struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Year      int32     `json:"year"`
	Quantity  int32     `json:"quantity"`
	Pounds    Pounds    `json:"weight_lb,omitempty"`
	Types     []string  `json:"types,omitempty"`
	CreatedAt time.Time `json:"-"`
}

func ValidateItem(v *validator.Validator, i *Item) {
	v.Check(i.Name != "", "name", "must be provided")
	v.Check(len(i.Name) <= 500, "name", "must not exceed 500 bytes")
	v.Check(i.Year >= 1900, "year", "must not precede 1900")
	v.Check(i.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	// TODO: determine if there's a way to specifically check for lack of user input
	v.Check(i.Quantity > 0, "quantity", "must be a positive integer")
	v.Check(i.Quantity <= 99, "quantity", "must be not exceed 99")
	v.Check(i.Pounds > 0, "pounds", "must be a postive integer")
	v.Check(i.Pounds <= 999, "pounds", "must not exceed 999")
	v.Check(i.Types != nil, "types", "must be provided")
	v.Check(len(i.Types) >= 1, "types", "must contain at least 1 type")
	v.Check(len(i.Types) <= 3, "types", "may contain at most 3 types")
	v.Check(validator.Unique(i.Types), "types", "must not contain duplicate types")
}
