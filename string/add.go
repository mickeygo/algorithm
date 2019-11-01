package string

import (
	"reflect"
	"strconv"
)

// Add two-sum string
func Add(a, b string) string {
	var l, s string
	if len(a) >= len(b) {
		l, s = a, b
	} else {
		l, s = b, a
	}

	var out string
	jw := 0
	step := len(l) - len(s)
	for i := len(l); i > 0; i = i - 1 {
		var x, y int
		x, _ = strconv.Atoi(reflect.ValueOf(l[i-1 : i]).String())
		if i > step {
			y, _ = strconv.Atoi(reflect.ValueOf(s[i-step-1 : i-step]).String())
		} else {
			y = 0
		}
		single := (int(x+y) + jw) % 10
		out = strconv.Itoa(single) + out
		jw = (x + y + jw) / 10
	}

	return out
}
