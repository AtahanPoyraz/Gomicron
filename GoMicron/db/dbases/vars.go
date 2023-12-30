package dbproc

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/AtahanPoyraz/cmd"
	"github.com/AtahanPoyraz/config"
	"github.com/AtahanPoyraz/db"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

var (
	msql     *db.DBInfo
	pqdb     *db.DBInfo
	conf     config.Config 
	async    sync.WaitGroup
	validate = validator.New()
	channel  = make(chan *db.DBResults, 1)
	l		 = log.New(os.Stdout, fmt.Sprintf("%sGOMICRON >> %s", cmd.TCYAN, cmd.TRESET), log.LstdFlags) 
)