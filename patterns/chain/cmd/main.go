package main

import "chain"

func main() {
	msg := chain.NewMessage(false, true, false)
	//создадим обработчики
	spamHandler := &chain.SpamHandler{}
	workHandler := &chain.WorkHandler{}
	advertHandler := &chain.AdvertHandler{}
	//установим им следующие обработчики
	spamHandler.SetNextHandler(workHandler)
	workHandler.SetNextHandler(advertHandler)

	spamHandler.Execute(msg)
}
