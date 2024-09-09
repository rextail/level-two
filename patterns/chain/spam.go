package chain

import "fmt"

type SpamHandler struct {
	next Handler
}

func (s *SpamHandler) SetNextHandler(h Handler) {
	s.next = h
}

func (s *SpamHandler) Execute(m Message) {
	if m.isSpam == true {
		fmt.Println("Помещаем письмо в папку спама")
		return
	}
	s.next.Execute(m)
}
