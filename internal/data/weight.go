package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Pounds int32

func (p *Pounds) UnmarshalJSON(v []byte) error {
	/*TODO: refactor so this handles two forms of input...
	  should accept "pounds": 5
	  should accept "pounds": "5 lbs"
	*/
	unquoted_v, err := strconv.Unquote(string(v))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	splits := strings.Split(unquoted_v, " ")
	if len(splits) != 2 || splits[1] != "lbs" {
		return ErrInvalidRuntimeFormat
	}
	i, err := strconv.ParseInt(splits[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*p = Pounds(i)
	return nil
}

func (p Pounds) MarshalJSON() ([]byte, error) {
	data := fmt.Sprintf("%d lbs", p)
	payload := strconv.Quote(data)
	return []byte(payload), nil
}
