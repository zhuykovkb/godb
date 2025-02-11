package server

import (
	"errors"
	"fmt"
	"goconcurrency/internal/compute"
	"goconcurrency/internal/config"
	"goconcurrency/internal/db"
	"goconcurrency/internal/logger"
	"goconcurrency/internal/semaphore"
	inmemory "goconcurrency/internal/storage/inMemory"
	"net"
	"time"
)

const ConnType = "tcp"

type Server struct {
	Listener net.Listener
	Db       *db.Db
	Config   *config.Config
}

func (s *Server) Run() {
	logger.Info("server started")
	defer logger.Info("server stopped")

	maxConn := uint32(s.Config.Network.MaxConnections)

	sem := semaphore.NewSemaphore(maxConn)
	logger.Info(fmt.Sprintf("max connections: %d", maxConn))

	for {
		sem.Acquire()

		conn, err := s.Listener.Accept()
		if err != nil {
			logger.Info("failed to accept connection")
			continue
		}

		go func(conn net.Conn) {
			defer func() {
				if err := recover(); err != nil {
					logger.Warn(fmt.Sprintf("recovered from panic err: %s", string(err.([]byte))))
				}
				sem.Release()
			}()
			s.handle(conn)
		}(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	idleTimeout := s.Config.Network.IdleTimeout
	logger.Info(fmt.Sprintf("idle timeout: %.2f", idleTimeout.Seconds()))

	maxMessageSize := s.Config.Network.GetMaxMessageSize()
	logger.Info(fmt.Sprintf("max message size: %d", maxMessageSize))

	buf := make([]byte, maxMessageSize)

	for {

		if idleTimeout.Seconds() > 0 {
			if err := conn.SetDeadline(time.Now().Add(idleTimeout)); err != nil {
				logger.Fatal(fmt.Sprintf("failed to set connection deadline: %s", err))
			}
		}

		n, err := conn.Read(buf)

		if err != nil {
			logger.Warn("failed to read from connection")
			break
		}

		if n > cap(buf) {
			logger.Warn("too many bytes read from connection")
			break
		}

		r, er := s.Db.HandleReq(string(buf[:n]))
		if er != nil {
			logger.Warn("failed to handle request")
			r = er.Error()
		}

		_, err = conn.Write([]byte(r))
		if err != nil {
			logger.Warn("failed to write response to connection")
		}
	}

}

func NewServer(cfg *config.Config) (*Server, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}

	database := db.GetDb(inmemory.NewEngine(), compute.NewParser())

	l, err := net.Listen(ConnType, cfg.Network.Address)
	if err != nil {
		return nil, err
	}

	return &Server{
		Listener: l,
		Db:       database,
		Config:   cfg,
	}, nil

}
