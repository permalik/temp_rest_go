package data

import (
	"time"
)

type Item struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Quantity  int32     `json:"quantity"`
	Pounds    Pounds    `json:"weight_lb,omitempty"`
	Types     []string  `json:"types,omitempty"`
	CreatedAt time.Time `json:"-"`
}
