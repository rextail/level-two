package chain

import "fmt"

type AdvertHandler struct {
	next Handler
}

func (s *AdvertHandler) SetNextHandler(h Handler) {
	s.next = h
}

func (s *AdvertHandler) Execute(m Message) {
	if m.isAdvert == true {
		fmt.Println("Помещаем письмо в папку промо")
		return
	}
	s.next.Execute(m)
}
