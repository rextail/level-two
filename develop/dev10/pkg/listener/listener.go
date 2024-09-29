package listener

import (
	"fmt"
	"io"
	"log"
	"net"
)

type TelnetListener struct {
	telnet net.Listener
}

func DefaultTelnetListener() (*TelnetListener, error) {
	const op = "listener.server.DefaultTelnetListener"
	listen, err := net.Listen("tcp", net.JoinHostPort("localhost", "23"))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &TelnetListener{
		telnet: listen,
	}, nil
}

func (t *TelnetListener) HandleConnections() {
	const op = "listener.listener.HandleConnections"
	for {
		conn, err := t.telnet.Accept()
		if err != nil {
			log.Printf("%s: listening to incoming connection failed: %v", op, err)
			continue
		}
		go func() {
			err = t.handleConnection(conn)
			if err != nil {
				log.Printf("%v", err)
			}
		}()
	}
}

func (t *TelnetListener) handleConnection(conn net.Conn) error {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("client has ended connection")
				break
			}
			return fmt.Errorf("error occurred reading from connection: %w", err)
		}
		_, err = conn.Write(buffer[:n])
		if err != nil {
			return fmt.Errorf("error occurred writing response: %w", err)
		}
	}

	return nil
}
