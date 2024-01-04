package auth

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/AtahanPoyraz/cmd"
	"github.com/AtahanPoyraz/config"
	"github.com/AtahanPoyraz/db"

	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
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
	conf     config.Config 
	pqdb     *db.DBInfo
	globalToken *Token
	async sync.WaitGroup
	channel  = make(chan *db.DBResults, 1)
	l		 = log.New(os.Stdout, fmt.Sprintf("%sGOMICRON >> %s", cmd.TCYAN, cmd.TRESET), log.LstdFlags) 
)