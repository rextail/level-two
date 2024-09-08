package facade

import "fmt"

type smartLight struct {
	brightness byte
}

func newSmartLight() *smartLight {
	return &smartLight{brightness: 0}
}

func (s *smartLight) TurnOffLight() {
	fmt.Println("Выключаю свет!")
	s.brightness = 0
}

func (s *smartLight) TurnNightMode() {
	fmt.Println("Включаю режим ночного освещения!")
	s.brightness = 15
}

func (s *smartLight) TurnOnLight() {
	fmt.Println("Включаю свет!")
	s.brightness = 80
}
