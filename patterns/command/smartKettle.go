package command

import "fmt"

type SmartKettle struct {
	turnedOn bool
}

func (s *SmartKettle) Enable() bool {
	if s.turnedOn {
		fmt.Println("Чайник уже включен, подождите, пожалуйста")
		return false
	}
	fmt.Println("Включаю чайник...")
	s.turnedOn = true
	return true
}

func (s *SmartKettle) Disable() bool {
	if !s.turnedOn {
		fmt.Println("Чайник уже выключен")
		return false
	}
	fmt.Println("Выключаю чайник...")
	s.turnedOn = false
	return true
}
