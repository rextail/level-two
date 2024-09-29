package main

import (
	"dev10/config"
	"dev10/internal/client"
	"dev10/pkg/listener"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	if err := cfg.Validate(); err != nil {
		log.Fatalf("incorrect startup parameters: %v", err)
	}

	telnetClient := client.New(*cfg)

	telnetListener, err := listener.DefaultTelnetListener()

	if err != nil {
		log.Fatalf("can't listen to default host and port: %v", err)
	}

	telnetClient.Connect()
	if err != nil {
		log.Fatalf("can't connect to the default listener: %v", err)
	}

	go func() {
		var input string
		for {
			fmt.Scan(&input)
			telnetClient.Send(input)
		}
	}()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go telnetListener.HandleConnections()

	go telnetClient.HandleResponses()

	<-sig

	log.Println("got signal to stop..")

	telnetClient.CloseConn()

}
