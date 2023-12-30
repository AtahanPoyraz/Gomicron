package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AtahanPoyraz/cmd"
	"github.com/AtahanPoyraz/db/dbases"
	protoc "github.com/AtahanPoyraz/protoc"
	"github.com/gorilla/mux"
)

type Books struct {
	l *log.Logger
}

func NewLibraryHandler(l *log.Logger) *Books {
	return &Books{l}
}

func (b *Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	b.l.Printf("%s[METHOD]%s  : %sGET%s >> %s\n", cmd.BGREEN_WHITE, cmd.TRESET, cmd.TGREEN, cmd.TRESET, r.URL.Path)
	books, err := dbproc.GETBooks()
	if err != nil {
		b.l.Printf("%s[ERROR]%s : %v", cmd.BRED_WHITE, cmd.TRESET, err)
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	for _, book := range books {
		b.l.Printf("%s[INFO]%s : %v", cmd.BGRAY_BLACK, cmd.TRESET, book)
	}
	response, _ := json.Marshal(books)
	w.Write(response)
}

func (b *Books) PostBooks(w http.ResponseWriter, r *http.Request) {
	b.l.Printf("%s[METHOD]%s  : %sPOST%s >> %s\n", cmd.BYELLOW_WHITE, cmd.TRESET, cmd.TYELLOW, cmd.TRESET, r.URL.Path)

	book := r.Context().Value(KeyStruct{}).(protoc.Book)
	dbproc.POSTBooks(&book)

	b.l.Printf("%s[INFO]%s : Book - %v\n", cmd.BGRAY_BLACK, cmd.TRESET, &book)
	response, _ := json.Marshal(book)
	w.Write(response)
}

func (b *Books) PutBooks(w http.ResponseWriter, r *http.Request) {
	b.l.Printf("%s[METHOD]%s  : %sPUT%s >> %s\n", cmd.BBLUE_WHITE, cmd.TRESET, cmd.TDBLUE, cmd.TRESET, r.URL.Path)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		b.l.Printf("%s[ERROR]%s : Unable to Convert ID: %v", cmd.BRED_WHITE, cmd.TRESET, err)
		return
	}

	book := r.Context().Value(KeyStruct{}).(protoc.Book)

	err = dbproc.PUTBooks(id, &book)
	if err == dbproc.ErrBookNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
		b.l.Printf("%s[ERROR]%s : Book Not Found: %v", cmd.BRED_WHITE, cmd.TRESET, err)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] : %v", err), http.StatusInternalServerError)
		b.l.Printf("%s[ERROR]%s : %v", cmd.BRED_WHITE, cmd.TRESET, err)
		return
	}
	book.ID = int64(id)
	b.l.Printf("%s[INFO]%s : book - %v\n", cmd.BGRAY_BLACK, cmd.TRESET, &book)
	response, _ := json.Marshal(book)
	w.Write(response)
}

func (b *Books) DeleteBooks(w http.ResponseWriter, r *http.Request) {
	b.l.Printf("%s[METHOD]%s  : %sDELETE%s >> %s\n", cmd.BORANGE_WHITE, cmd.TRESET, cmd.TORANGE, cmd.TRESET, r.URL.Path)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		b.l.Printf("%s[ERROR]%s : Unable to Convert ID: %v", cmd.BRED_WHITE, cmd.TRESET, err)
		return
	}

	book := r.Context().Value(KeyStruct{}).(protoc.Book)

	err = dbproc.DELETEBooks(id, &book)
	if err == dbproc.ErrBookNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
		b.l.Printf("%s[ERROR]%s : Book Not Found: %v", cmd.BRED_WHITE, cmd.TRESET, err)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] : %v", err), http.StatusInternalServerError)
		b.l.Printf("%s[ERROR]%s : %v", cmd.BRED_WHITE, cmd.TRESET, err)
		return
	}
	b.l.Printf("%s[INFO]%s : Book Delete Process has been Successful\n", cmd.BGRAY_BLACK, cmd.TRESET)
	response, _ := json.Marshal(" [INFO] : Book Delete Process has been Successful")
	w.Write(response)
}

func (b *Books) MiddlewareBooksValidation(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        var book protoc.Book
        if r.Body == nil {
            http.Error(rw, "Request body is empty", http.StatusBadRequest)
            return
        }
        err := json.NewDecoder(r.Body).Decode(&book)
        if err != nil {
            b.l.Printf("%s[ERROR]%s : Deserializing Data: %v", cmd.BRED_WHITE, cmd.TRESET, err)
            http.Error(rw, "Error reading data", http.StatusBadRequest)
            return
        }

        ctx := context.WithValue(r.Context(), KeyStruct{}, book)
        req := r.WithContext(ctx)

        next.ServeHTTP(rw, req)
    })
}

