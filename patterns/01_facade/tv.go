package facade

import "fmt"

type smartTv struct {
	channels byte
}

func newSmartTv() *smartTv {
	return &smartTv{channels: 255}
}

func (s *smartTv) WatchMovie(title string) {
	fmt.Printf("Включаю фильм %s, приятного просмотра!\n", title)
}

func (s *smartTv) WatchChannel(number int) {
	if number < 0 || number > int(s.channels) {
		fmt.Printf("Я не могу включить такой канал %d, может посмотрим другой?\n", number)
	} else {
		fmt.Printf("Включаю канал %d, приятного просмотра!\n", number)
	}
}
