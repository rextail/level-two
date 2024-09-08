package facade

import "fmt"

type smartKettle struct {
	temp       byte
	waterLevel byte
}

func newSmartKettle() *smartKettle {
	return &smartKettle{
		temp:       24,
		waterLevel: 50,
	}
}

func (s *smartKettle) HeatWater() {
	if s.temp >= 80 {
		fmt.Println("Чайник уже горячий!")
	}
	if s.waterLevel >= 20 {
		fmt.Println("Сейчас вскипячу вам водичку...")
	} else {
		fmt.Println("В чайнике мало воды!")
	}
}
