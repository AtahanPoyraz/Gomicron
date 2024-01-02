package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AtahanPoyraz/cmd"
	dbproc "github.com/AtahanPoyraz/db/dbases"
	protoc "github.com/AtahanPoyraz/protoc"
	"github.com/gorilla/mux"
)

type Students struct {
	l *log.Logger
}

func NewStudentHandler(l *log.Logger) *Students {
	return &Students{l}
}

func (s *Students) GetStudents (w http.ResponseWriter, r *http.Request) {
	s.l.Printf("%s[METHOD]%s  : %sGET%s >> %s\n",cmd.BGREEN_WHITE, cmd.TRESET, cmd.TGREEN, cmd.TRESET, r.URL.Path)
	students, err := dbproc.GETStudents()
	if err != nil {
		s.l.Printf("%s[ERROR]%s : %v",cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(0)
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	for _, student := range students {
		s.l.Printf("%s[INFO]%s : %v", cmd.BGRAY_BLACK, cmd.TRESET, student)
	}
	response, _ := json.Marshal(students)
	w.Write(response)
}

func (s *Students) PostStudents (w http.ResponseWriter, r *http.Request) {
	s.l.Printf("%s[METHOD]%s  : %sPOST%s >> %s\n",cmd.BYELLOW_WHITE, cmd.TRESET, cmd.TYELLOW, cmd.TRESET, r.URL.Path)

	std := r.Context().Value(KeyStruct{}).(protoc.Student)
	dbproc.POSTStudents(&std)

	s.l.Printf("%s[INFO]%s : Student - %v\n", cmd.BGRAY_BLACK, cmd.TRESET, &std)
	response, _ := json.Marshal(std)
	w.Write(response)
}

func (s *Students) PutStudents (w http.ResponseWriter, r *http.Request) {
	s.l.Printf("%s[METHOD]%s  : %sPUT%s >> %s\n",cmd.BBLUE_WHITE, cmd.TRESET, cmd.TDBLUE, cmd.TRESET, r.URL.Path)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		s.l.Printf("%s[ERROR]%s : Unable to Convert ID: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) //Geri Bildirim
		return
	}

	std := r.Context().Value(KeyStruct{}).(protoc.Student)

	err = dbproc.PUTStudents(id, &std)
	if err == dbproc.ErrStudentNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		s.l.Printf("%s[ERROR]%s : Product Not Found: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) 

		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] : %s", err), http.StatusInternalServerError)
		s.l.Printf("%s[ERROR]%s : %v",cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(0)

		return
	}
	std.Id = int64(id)
	s.l.Printf("%s[INFO]%s : Student - %v\n", cmd.BGRAY_BLACK, cmd.TRESET, &std)
	response, _ := json.Marshal(std)
	w.Write(response)
}

func (s *Students) DeleteStudents (w http.ResponseWriter, r *http.Request) {
	s.l.Printf("%s[METHOD]%s  : %sDELETE%s >> %s\n",cmd.BORANGE_WHITE, cmd.TRESET, cmd.TORANGE, cmd.TRESET, r.URL.Path)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		s.l.Printf("%s[ERROR]%s : Unable to Convert ID: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) //Geri Bildirim
		return
	}

	std := r.Context().Value(KeyStruct{}).(protoc.Student)

	err = dbproc.DELETEStudents(id, &std)
	if err == dbproc.ErrStudentNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		s.l.Printf("%s[ERROR]%s : Product Not Found: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) 

		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] : %s", err), http.StatusInternalServerError)
		s.l.Printf("%s[ERROR]%s : %v",cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(0)

		return
	}
	s.l.Printf("%s[INFO]%s : Student Delete Proccess been Succsessful\n", cmd.BGRAY_BLACK, cmd.TRESET)
	response, _ := json.Marshal(" [INFO] : Student Delete Proccess been Succsessful")
	w.Write(response)
}

//Kullanıcının attıgı json datayı burada alacagız.
func (s *Students) MiddlewareStudentValidation(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var std protoc.Student
		if r.Body == nil {
            http.Error(rw, "Request body is empty", http.StatusBadRequest)
            return
        }
        err := json.NewDecoder(r.Body).Decode(&std)
        if err != nil {
            s.l.Printf("%s[ERROR]%s : Deserializing Data: %v", cmd.BRED_WHITE, cmd.TRESET , err)
            http.Error(rw, "Error reading data", http.StatusBadRequest)
            return
        }

        ctx := context.WithValue(r.Context(), KeyStruct{}, std)
        req := r.WithContext(ctx)

        next.ServeHTTP(rw, req)
    })
}
