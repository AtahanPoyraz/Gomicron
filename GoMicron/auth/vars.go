package auth

import (
	"log"
)

type Server struct {
	l *log.Logger
}

func ServerHandler(l *log.Logger) *Server {
	return &Server{l}
}