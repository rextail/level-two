package dev01

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

/*
	Создать программу, печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
	Использовать библиотеку github.com/beevik/ntp. Написать программу, печатающую текущее время / точное время с
	использованием этой библиотеки.
*/

func CurrentTime() {
	const op = `develop.dev01.CurrentTime`
	options := ntp.QueryOptions{
		Timeout: 20 * time.Second,
	}
	resp, err := ntp.QueryWithOptions("0.beevik-ntp.pool.ntp.org", options)
	if err != nil {
		log.Fatalf("%s %v", op, err)
	}
	err = resp.Validate()
	if err != nil {
		log.Fatalf("%s %v", op, err)
	}
	srvTime := time.Now().Add(resp.ClockOffset).Format(time.TimeOnly)
	fmt.Println(srvTime)
}
