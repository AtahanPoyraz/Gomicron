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

	_ "github.com/go-sql-driver/mysql"
)


func init() {
	c, err := config.ReadConfigFromFile("./config.yml")
	if err != nil {
		l.Printf("%s[ERROR]%s : Config file read operation failed: %v", cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(1)
	}
	conf = c
	msql = db.NewDB(conf.Services.MySql.DBHost, conf.Services.MySql.DBPort, conf.Services.MySql.DBUser,
		conf.Services.MySql.DBPass, conf.Services.MySql.DBName, conf.Services.MySql.DBSSLMode,
		conf.Services.MySql.DBNET, conf.Services.MySql.DBTimeout)
}

type Books []*protoc.Book

func (b *Books) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}

func (b *Books) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

func GETBooks() (Books, error) {
	async.Add(1)
	go msql.OpenDB(conf.Services.MySql.DBSrvname, channel)

	result := <-channel

	if result.Err != nil {
		return nil, fmt.Errorf("error: %v", result.Err)
	}

	db := result.DB

	rows, err := db.Query("SELECT ID, BOOK_NAME, BOOK_NUM, BOOK_WRITER FROM BOOKS")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var bookList Books
	for rows.Next() {
		var book protoc.Book
		if err := rows.Scan(&book.ID, &book.BOOK_NAME, &book.BOOK_NUM, &book.BOOK_WRITER); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		bookList = append(bookList, &book)
	}

	return bookList, nil
}

func POSTBooks(b *protoc.Book) error {
	if err := validate.Struct(b); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	async.Add(1)
	go msql.OpenDB(conf.Services.MySql.DBSrvname, channel)

	result := <-channel

	if result.Err != nil {
		return fmt.Errorf("error: %v", result.Err)
	}

	db := result.DB

	b.ID = int64(GetBooksID())

	_, err := db.Exec("INSERT INTO BOOKS(ID, BOOK_NAME, BOOK_NUM, BOOK_WRITER) VALUES(?, ?, ?, ?);", b.ID, b.BOOK_NAME, b.BOOK_NUM, b.BOOK_WRITER)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

func PUTBooks(id int, b *protoc.Book) error {
	if err := validate.Struct(b); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	async.Add(1)
	go msql.OpenDB(conf.Services.MySql.DBSrvname, channel)
	result := <-channel
	if result.Err != nil {
		return fmt.Errorf("error: %v", result.Err)
	}
	db := result.DB

	existingBook, _, err := FindBook(id)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	existingBook.BOOK_NAME = b.BOOK_NAME
	existingBook.BOOK_NUM = b.BOOK_NUM
	existingBook.BOOK_WRITER = b.BOOK_WRITER

	_, err = db.Exec("UPDATE BOOKS SET BOOK_NAME = ?, BOOK_NUM = ?, BOOK_WRITER = ? WHERE ID = ?;",
		existingBook.BOOK_NAME, existingBook.BOOK_NUM, existingBook.BOOK_WRITER, id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

func DELETEBooks(id int, b *protoc.Book) error {
	if err := validate.Struct(b); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	async.Add(1)
	go msql.OpenDB(conf.Services.MySql.DBSrvname, channel)
	result := <-channel
	if result.Err != nil {
		return fmt.Errorf("error: %v", result.Err)
	}

	db := result.DB

	_, err := db.Exec("DELETE FROM BOOKS WHERE ID = ?;", id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	return nil
}

func GetBooksID() int {
	bookL, _ := GETBooks()

	if len(bookL) == 0 {
		return 1
	}

	lbook := bookL[len(bookL)-1]
	return int(lbook.ID) + 1
}

var ErrBookNotFound = fmt.Errorf("book not found")

func FindBook(id int) (*protoc.Book, int, error) {
	bookList, err := GETBooks()
	if err != nil {
		return nil, -1, err
	}
	for i, b := range bookList {
		if int(b.ID) == id {
			return b, i, nil
		}
	}

	return nil, -1, ErrBookNotFound
}
