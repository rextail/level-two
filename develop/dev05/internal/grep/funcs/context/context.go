package context

import (
	"dev05/internal/grep/funcs"
)

func New() funcs.RangeFunc {
	return func(index, count int) [2]int {
		rng := [2]int{0, 0}

		rng[0] = max(0, index-count)

		rng[1] = index + count

		return rng
	}
}
