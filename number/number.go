package number

import (
	"bytes"
	"strconv"
)

type Number struct {
	Str string
	Val int
}

func (n Number) MarshalJSON() (bs []byte, err error) {
	if n.Val == 0 && n.Str == "" {
		bs = []byte("null")
		return
	}

	bs = []byte(strconv.Itoa(n.Val))
	return
}

func (n *Number) UnmarshalJSON(bs []byte) (err error) {
	l := len(bs)
	if l == 4 && bytes.Equal(bs, []byte("null")) {
		return
	}

	if l >= 2 && bs[0] == '"' && bs[l-1] == '"' {
		bs = bs[1 : l-1]
	}

	n.Str = string(bs)
	n.Val, err = strconv.Atoi(n.Str)
	return
}
