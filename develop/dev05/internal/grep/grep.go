package grep

import (
	"dev05/config"
	"dev05/internal/grep/funcs"
	"dev05/internal/grep/funcs/after"
	"dev05/internal/grep/funcs/before"
	"dev05/internal/grep/funcs/context"
)

type Grepper struct {
	count int
	fn    funcs.RangeFunc
}

func New(cfg config.Config) *Grepper {
	var grp Grepper
	switch {
	case cfg.After > 0:
		grp.fn = after.New()
		grp.count = cfg.After
	case cfg.Before > 0:
		grp.fn = before.New()
		grp.count = cfg.Before
	case cfg.Context > 0:
		grp.fn = context.New()
		grp.count = cfg.Context
	}
	return &grp
}

// GrepABC возвращает срез диапазонов, которые необходимо включить в ответ. Функция использует номера строк, в которых
// встречается заданное выражение и включает в диапазон следующие и предыдущие count строк. Если один диапазон включает
// в себя другой, то они объединяются. Применяется для флагов -A, -B или -C
func (g *Grepper) GrepABC(indexes []int) [][2]int {
	var ranges [][2]int

	for _, index := range indexes {

		rng := g.fn(index, g.count)

		if len(ranges) == 0 {
			ranges = append(ranges, rng)

		} else {
			//Если в диапазонах уже есть хоть что-то, начинаем проверять на пересечения
			oldRng := ranges[len(ranges)-1]
			if oldRng[1] >= rng[0] {
				//Если правая граница старого диапазона больше, чем левая граница нового диапазона, тогда
				//объединим диапазоны, обновив правую границу первого
				ranges[len(ranges)-1][1] = rng[1]
			} else {
				//Если диапазоны не пересекаются, добавим в список
				ranges = append(ranges, rng)
			}
		}
	}

	return ranges
}
