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

type Token struct {
	AuthToken string
}

var (
	globalToken *Token
)