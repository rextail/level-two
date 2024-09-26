package pkg

import (
	"errors"
)

const (
	// Buffer ...
	Buffer = 2<<16 - 1
	// Disconnect ...
	Disconnect = "~!"
)

type Netcater struct {
	host     string
	port     string
	protocol string
}

func NewNetcater(host, port, protocol string) *Netcater {
	return &Netcater{
		host:     host,
		port:     port,
		protocol: protocol,
	}
}

func (nc *Netcater) Start() error {
	switch nc.protocol {
	case "tcp":
		if err := nc.TCPConnector(); err != nil {
			return err
		}
	case "udp":
		if err := nc.UDPConnector(); err != nil {
			return err
		}
	default:
		return errors.New("Netcater: provided protocol usuported")
	}
	return nil
}
