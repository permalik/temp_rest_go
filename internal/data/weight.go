package data

import (
	"fmt"
	"strconv"
)

type Pounds int32

func (p Pounds) MarshalJSON() ([]byte, error) {
	data := fmt.Sprintf("%d lbs", p)
	payload := strconv.Quote(data)
	return []byte(payload), nil
}
