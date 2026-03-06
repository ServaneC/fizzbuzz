package fizzbuzz

import (
	"strconv"
)

func Generate(p Params) []string {
	ret := make([]string, p.Limit)

	for i := range ret {
		if (i+1)%p.Int1 == 0 {
			ret[i] = p.Str1
		}
		if (i+1)%p.Int2 == 0 {
			ret[i] += p.Str2
		}
		if ret[i] == "" {
			ret[i] = strconv.Itoa(i + 1)
		}
	}
	return ret
}
