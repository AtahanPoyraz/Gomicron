package dbproc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/AtahanPoyraz/cmd"
	"github.com/AtahanPoyraz/config"
	"github.com/AtahanPoyraz/db"
	protoc "github.com/AtahanPoyraz/protoc"

	_ "github.com/lib/pq"
)


func init() {
	c, err := config.ReadConfigFromFile("./config.yml")
	if err != nil {
		l.Printf("%s[ERROR]%s : Read operation failed please check file : %v", cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(0)
	}
	conf = c
	pqdb = db.NewDB(conf.Services.Postgres.DBHost, conf.Services.Postgres.DBPort, conf.Services.Postgres.DBUser,
					conf.Services.Postgres.DBPass, conf.Services.Postgres.DBName, conf.Services.Postgres.DBSSLMode,
	   				conf.Services.Postgres.DBNET, conf.Services.Postgres.DBTimeout)
}

type Students []*protoc.Student

func (s *Students) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

func (s *Students) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func GETStudents() (Students, error) {
	async.Add(1)
	go pqdb.OpenDB(conf.Services.Postgres.DBSrvname, channel)

	result := <- channel

	if result.Err != nil {
		return nil, fmt.Errorf("error: %v", result.Err)
	}

	db := result.DB

	rows, err := db.Query("SELECT ID, NAME, SURNAME, NOTE FROM OGRENCILER")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var stdList Students
	for rows.Next() {
		var std protoc.Student
		err := rows.Scan(&std.Id, &std.Name, &std.Surname, &std.Note)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		stdList = append(stdList, &std)
	} 
	
	return stdList, nil
}

func POSTStudents(s *protoc.Student) error {
	valErr := validate.Struct(s)
	if valErr != nil {
		return fmt.Errorf("validation error: %v", valErr)
	}

	async.Add(1)
	go pqdb.OpenDB(conf.Services.Postgres.DBSrvname, channel)

	result := <-channel

	if result.Err != nil {
		return fmt.Errorf("error: %v", result.Err)
	}

	db := result.DB

	s.Id = int64(GETID())

	_, err := db.Exec("INSERT INTO OGRENCILER(ID, NAME, SURNAME, NOTE) VALUES($1, $2, $3, $4);", s.Id, s.Name, s.Surname, s.Note)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

func PUTStudents(id int, s *protoc.Student) error {
	// Validate the input student
	valErr := validate.Struct(s)
	if valErr != nil {
		return fmt.Errorf("validation error: %v", valErr)
	}

	// Open database connection
	async.Add(1)
	go pqdb.OpenDB(conf.Services.Postgres.DBSrvname, channel)
	result := <-channel
	if result.Err != nil {
		return fmt.Errorf("error: %v", result.Err)
	}
	db := result.DB

	// Find the student with the given ID
	existingStudent, _, err := FindStd(id)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	// Update the fields of the existing student with the new values
	existingStudent.Name = s.Name
	existingStudent.Surname = s.Surname
	existingStudent.Note = s.Note

	// Execute the update query
	_, err = db.Exec("UPDATE OGRENCILER SET NAME = $2, SURNAME = $3, NOTE = $4 WHERE ID = $1;",
		id, existingStudent.Name, existingStudent.Surname, existingStudent.Note)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

func DELETEStudents(id int, s *protoc.Student) error {
	valErr := validate.Struct(s)
	if valErr != nil {
		return fmt.Errorf("validation error: %v", valErr)
	}
	async.Add(1)
	go pqdb.OpenDB(conf.Services.Postgres.DBSrvname, channel)
	result := <-channel
	if result.Err != nil {
		return fmt.Errorf("error: %v", result.Err)
	}

	db := result.DB

	_, err := db.Exec("DELETE FROM OGRENCILER WHERE ID = $1;", id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	return nil
}

func GETID() int {
    stdL, _ := GETStudents()

    if len(stdL) == 0 {
        return 1 
    }

    lstudent := stdL[len(stdL)-1]
    return int(lstudent.Id) + 1
}

var ErrStudentNotFound = fmt.Errorf("product not found")

func FindStd(id int) (*protoc.Student, int, error) {
	stdList, err := GETStudents()
	if err != nil {
		return nil, -1, err
	}
	for i, s := range stdList {
		if int(s.Id) == id {
			return s, i, nil
		}
	}

	return nil, -1, ErrStudentNotFound
}