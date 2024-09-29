package client

import (
	"dev10/config"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Client struct {
	addr    string
	timeout time.Duration
	conn    net.Conn
}

func New(cfg config.Config) *Client {
	timeout, _ := time.ParseDuration(cfg.Timeout)
	return &Client{
		addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		timeout: timeout,
	}
}

func (c *Client) Connect() error {
	const op = "client.Connect"
	conn, err := net.DialTimeout("tcp", c.addr, c.timeout)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	c.conn = conn
	log.Printf("%s: connection to %s was established!\n", op, c.addr)
	return nil
}

func (c *Client) Send(msg string) error {
	const op = "client.Send"

	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (c *Client) HandleResponses() {
	const op = "client.HandleResponse"
	buffer := make([]byte, 1024)
	for {
		_, err := c.conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("connection was ended")
				break
			}
			log.Printf("%s: failed reading response: %v", op, err)
		} else {
			fmt.Println(string(buffer))
		}
	}
}

func (c *Client) CloseConn() {
	if c.conn != nil {
		c.conn.Close()
	}
}
