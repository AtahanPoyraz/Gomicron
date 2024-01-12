package router

import (
	"log"
	"net/http"

	"github.com/AtahanPoyraz/auth"
	"github.com/AtahanPoyraz/handlers"

	"github.com/gorilla/mux"
)

func SetRoutes(r *mux.Router, l *log.Logger) {

	sh := handlers.NewStudentHandler(l) 
	ph := handlers.NewProductsHandler(l)   
	lh := handlers.NewLibraryHandler(l)                                            

//---[ SERVICE PATHS ]---------------------------------------------------------------------------------------------------------------------------------//

	//---[ SERVICE 1]---------------------------------------------------------------------------------------------------------------------------------//

	SgetRouter := r.Methods(http.MethodGet).Subrouter()
	SgetRouter.HandleFunc("/students/get/", sh.GetStudents)
	SgetRouter.Use(auth.HandleAuthMiddleware)

	SpostRouther := r.Methods(http.MethodPost).Subrouter()
	SpostRouther.HandleFunc("/students/post/", sh.PostStudents)
	SpostRouther.Use(auth.HandleAuthMiddleware)
	SpostRouther.Use(sh.MiddlewareStudentValidation)

	SputRouther := r.Methods(http.MethodPut).Subrouter()
	SputRouther.HandleFunc("/students/put/{id:[0-9]+}", sh.PutStudents)
	SputRouther.Use(auth.HandleAuthMiddleware)
	SputRouther.Use(sh.MiddlewareStudentValidation)

	SdeleteRouther := r.Methods(http.MethodDelete).Subrouter()
	SdeleteRouther.HandleFunc("/students/delete/{id:[0-9]+}", sh.DeleteStudents)
	SdeleteRouther.Use(auth.HandleAuthMiddleware)
	SdeleteRouther.Use(sh.MiddlewareStudentValidation)

	//---[ SERVICE 2]---------------------------------------------------------------------------------------------------------------------------------//

	PgetRouther := r.Methods(http.MethodGet).Subrouter()
	PgetRouther.HandleFunc("/products/get/", ph.GetProduct)

	PpostRouther := r.Methods(http.MethodPost).Subrouter()
	PpostRouther.HandleFunc("/products/post/", ph.AddProduct)
	PpostRouther.Use(ph.MiddlewareProductValidation)

	PputRouther := r.Methods(http.MethodPut).Subrouter()
	PputRouther.HandleFunc("/products/put/{id:[0-9]+}", ph.UpdateProduct)
	PputRouther.Use(ph.MiddlewareProductValidation)

	PdeleteRouther := r.Methods(http.MethodDelete).Subrouter()
	PdeleteRouther.HandleFunc("/products/delete/{id:[0-9]+}", ph.DeleteProduct)
	PdeleteRouther.Use(ph.MiddlewareProductValidation)

	//---[ SERVICE 3]---------------------------------------------------------------------------------------------------------------------------------//

	LgetRouter := r.Methods(http.MethodGet).Subrouter()
	LgetRouter.HandleFunc("/books/get/", lh.GetBooks)

	LpostRouther := r.Methods(http.MethodPost).Subrouter()
	LpostRouther.HandleFunc("/books/post/", lh.PostBooks)
	LpostRouther.Use(lh.MiddlewareBooksValidation)

	LputRouther := r.Methods(http.MethodPut).Subrouter()
	LputRouther.HandleFunc("/books/put/{id:[0-9]+}", lh.PutBooks)
	LputRouther.Use(lh.MiddlewareBooksValidation)

	LdeleteRouther := r.Methods(http.MethodDelete).Subrouter()
	LdeleteRouther.HandleFunc("/books/delete/{id:[0-9]+}", lh.DeleteBooks)
	LdeleteRouther.Use(lh.MiddlewareBooksValidation)

}