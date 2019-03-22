package strut

import (
	"strings"
)

// возвращает позицию разделите d в строке который встречается первым
func IndexOfFirst(s string, d []string) int {
	var (
		p, b int
	)
	p = -1
	for _, val := range d {
		b = strings.Index(s, val)
		if b > -1 && ((b < p && p > -1) || p == -1) {
			p = b
		}
	}
	return p
} // end IndexOfFirst
