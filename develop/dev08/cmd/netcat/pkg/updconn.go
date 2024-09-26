package pkg

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func (nc *Netcater) UDPConnector() error {
	creds := nc.host + ":" + nc.port
	addr, err := net.ResolveUDPAddr("udp", creds)
	if err != nil {
		return errors.New("Netcater: can not resolve UDP address")
	}
	conn, err := net.DialUDP(nc.protocol, nil, addr)
	if err != nil {
		return errors.New("Netcater: can not build UDP connection")
	}
	log.Println("Netcater: connected to", creds)
	nc.transferUDP(conn)
	return nil
}

type udpProgress struct {
	remoteAddr net.Addr
	bytes      uint64
}

func (nc *Netcater) transferUDP(conn net.Conn) {
	c := make(chan udpProgress)

	copy := func(r io.ReadCloser, w io.WriteCloser, ra net.Addr) {
		defer func() {
			r.Close()
			w.Close()
		}()

		buf := make([]byte, Buffer)
		bytes := uint64(0)
		var n int
		var err error
		var addr net.Addr

		for {
			if connUDP, ok := r.(*net.UDPConn); ok {
				n, addr, err = connUDP.ReadFrom(buf)
				if connUDP.RemoteAddr() == nil && ra == nil {
					ra = addr
					c <- udpProgress{remoteAddr: ra}
				}
			} else {
				n, err = r.Read(buf)
			}
			if err != nil {
				if err != io.EOF {
					log.Printf("[%s]: ERROR: %s\n", ra, err)
				}
				break
			}
			if string(buf[0:n-1]) == Disconnect {
				break
			}

			if con, ok := w.(*net.UDPConn); ok && con.RemoteAddr() == nil {
				n, err = con.WriteTo(buf[0:n], ra)
			} else {
				n, err = w.Write(buf[0:n])
			}
			if err != nil {
				log.Printf("[%s]: ERROR: %s\n", ra, err)
				break
			}
			bytes += uint64(n)
		}
		c <- udpProgress{bytes: bytes}
	}

	ra := conn.RemoteAddr()
	go copy(conn, os.Stdout, ra)
	if ra == nil {
		p := <-c
		ra = p.remoteAddr
		fmt.Println(fmt.Sprintf("[%s]: Datagram has been received\n", ra))
	}
	go copy(os.Stdin, conn, ra)

	p := <-c
	fmt.Println(fmt.Sprintf(
		"[%s]: Connection has been closed, %d bytes has been received\n",
		ra,
		p.bytes,
	))
	p = <-c
	fmt.Println(fmt.Sprintf(
		"[%s]: Local peer has been stopped, %d bytes has been sent\n",
		ra,
		p.bytes,
	))
}
