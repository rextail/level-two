package chain

import "fmt"

type WorkHandler struct {
	next Handler
}

func (s *WorkHandler) SetNextHandler(h Handler) {
	s.next = h
}

func (s *WorkHandler) Execute(m Message) {
	if m.isWork == true {
		fmt.Println("Помещаем письмо в рабочую папку")
		return
	}
	s.next.Execute(m)
}
