package facade

import "fmt"

type smartJalousie struct {
	open bool
}

func newSmartJalousie() *smartJalousie {
	return &smartJalousie{
		open: false,
	}
}
func (s *smartJalousie) Open() {
	if !s.open {
		fmt.Println("Открываю занавески...")
	} else {
		fmt.Println("Датчики говорят, что занавески открыты")
	}
}

func (s *smartJalousie) Close() {
	if s.open {
		fmt.Println("Закрываю занавески")
	} else {
		fmt.Println("Датчики говорят, что занавески уже закрыты")
	}
}
