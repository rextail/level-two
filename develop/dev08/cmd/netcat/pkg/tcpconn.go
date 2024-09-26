package pkg

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

func (nc *Netcater) TCPConnector() error {
	creds := nc.host + ":" + nc.port
	conn, err := net.Dial("tcp", creds)
	if err != nil {
		return errors.New("Netcater: can not build TCP connection")
	}
	println(fmt.Sprintln("Netcater: connected to", creds))
	nc.transferTCP(conn)
	return nil
}

type tcpProgress struct {
	bytes uint64
}

func (nc *Netcater) transferTCP(conn net.Conn) {
	c := make(chan tcpProgress)

	copy := func(r io.ReadCloser, w io.WriteCloser) {
		defer func() {
			r.Close()
			w.Close()
		}()
		n, err := io.Copy(w, r)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s]: ERROR: %s\n", conn.RemoteAddr(), err))
		}
		c <- tcpProgress{bytes: uint64(n)}
	}

	go copy(conn, os.Stdout)
	go copy(os.Stdin, conn)

	p := <-c
	fmt.Println(fmt.Sprintf(
		"[%s]: Connection has been closed by remote peer, %d bytes has been received\n",
		conn.RemoteAddr(),
		p.bytes,
	))
	p = <-c
	fmt.Println(fmt.Sprintf(
		"[%s]: Local peer has been stopped, %d bytes has been sent\n",
		conn.RemoteAddr(),
		p.bytes,
	))
}
