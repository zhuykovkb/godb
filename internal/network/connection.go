package network

import (
	"goconcurrency/internal/logger"
	"io"
	"net"
)

const TcpProto = "tcp"

type Connection struct {
	c             net.Conn
	maxBufferSize uint
}

func newConnection(network string, address string, maxBufferSize uint) (*Connection, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &Connection{
		c:             c,
		maxBufferSize: maxBufferSize,
	}, err
}

func (c *Connection) Close() {
	_ = c.c.Close()
}

func (c *Connection) Receive(b []byte) (n int, err error) {
	return 0, nil
}

func (c *Connection) Send(b []byte) ([]byte, error) {
	_, err := c.c.Write(b)
	if err != nil {
		return nil, err
	}

	resp := make([]byte, c.maxBufferSize)
	bytesRead, err := c.c.Read(resp)

	if err != nil && err != io.EOF {
		logger.Fatal(err.Error())
	} else if bytesRead >= int(c.maxBufferSize) {
		logger.Fatal("max buffer size exceeded")
	}

	return resp[:bytesRead], nil
}

func NewTcpConnection(address string, maxBufferSize uint) (*Connection, error) {
	c, err := newConnection(TcpProto, address, maxBufferSize)
	if err != nil {
		return &Connection{}, err
	}
	return c, nil
}
