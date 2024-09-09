package chain

// Максимально упростим задачу определения типа сообщения
type Message struct {
	msg      string
	isSpam   bool
	isWork   bool
	isAdvert bool
}

func NewMessage(isSpam, isWork, isAdvert bool) Message {
	return Message{
		"Добрый день! Как насчет присоединиться к команде Wildberries?",
		isSpam,
		isWork,
		isAdvert,
	}
}
