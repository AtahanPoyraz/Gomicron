package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

//---[ DATABASE STRUCT ]---------------------------------------------------------------------------------------------------------------------------------------------------//

type DBInfo struct {
	Host 	 string
	Port 	 int
	User 	 string
	Pass	 string
	Name 	 string
	SSLmode  string
	Net      string
	Timeout  int
}

type DBResults struct {
	DB  *sql.DB
	Err error
}

//---[ DATABASE OPEN FUNC ]-------------------------------------------------------------------------------------------------------------------------------------------------//

func (db DBInfo) OpenDB(driver string, channel chan *DBResults) {
	switch {
		case driver == "postgres":
			var dbInfo = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d", db.Host, db.Port, db.Name, db.User, db.Pass, db.SSLmode, db.Timeout)
			dbase, err := sql.Open(driver, dbInfo)
			if err != nil {
				channel <- &DBResults{DB: nil, Err: err}
				return
			}		
			channel <- &DBResults{DB: dbase, Err: nil}

		case driver == "mysql":
			var dbInfo = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", db.User, db.Pass, db.Net, db.Host, db.Port, db.Name)
			dbase, err := sql.Open(driver, dbInfo)
			if err != nil {
				channel <- &DBResults{DB: nil, Err: err}
				return
			}
			channel <- &DBResults{DB: dbase, Err: nil}
		}	
} 

//---[ SUMMON DATABASE FUNC ]-----------------------------------------------------------------------------------------------------------------------------------------------//

func NewDB(host string, port int, user, password, name, sslmode, net string, timeout int) *DBInfo {
	return &DBInfo{
		Host: host,
		Port: port,
		User: user,
		Pass: password,
		Name: name,
		SSLmode: sslmode,
		Net: net,
		Timeout: timeout,
	}
}