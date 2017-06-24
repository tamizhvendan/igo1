package jwt

import (
	"strconv"
	"strings"
)

type Sub uint

func (s *Sub) UnmarshalJSON(b []byte) error {
	sub := strings.Replace(string(b), `"`, ``, 2)
	v, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return err
	}
	*s = Sub(uint(v))
	return nil
}
