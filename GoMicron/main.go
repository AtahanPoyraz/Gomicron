package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AtahanPoyraz/auth"
	"github.com/AtahanPoyraz/cmd"
	"github.com/AtahanPoyraz/config"
	"github.com/AtahanPoyraz/handlers"

	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//---[ STARTUP ]---------------------------------------------------------------------------------------------------------------------------------------//

func main() {
	r := mux.NewRouter()
	l := log.New(os.Stdout, fmt.Sprintf("%sGOMICRON-API >> %s", cmd.TCYAN, cmd.TRESET), log.LstdFlags) 

	config, err := config.ReadConfigFromFile("./config.yml")
	
	if err != nil {
		l.Printf("%s[ERROR]%s : Read operation failed please check file : (%e)", cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(1)
	}

//---[ PAGE ]-----------------------------------------------------------------------------------------------------------------------------------------//

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates/assets/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/html/main.html")
	})

	// This route handles any path and applies the AuthMiddleware
	r.HandleFunc("/dashboard/", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/html/dashboard.html")
	})))
	
	// The generic path prefix route, should come after more specific routes
	//r.PathPrefix("/{.*}").Handler(auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "./templates/html/main.html")
	//})))

//---[ SERVER PATHS ]---------------------------------------------------------------------------------------------------------------------------------//
	
	srvh := auth.ServerHandler(l)

	authRouter := r.Methods(http.MethodPost).Subrouter()
	authRouter.HandleFunc("/gomicron/backend/user/auth/", srvh.AuthUser)
	
//---[ HANDLERS ]--------------------------------------------------------------------------------------------------------------------------------------//

	sh := handlers.NewStudentHandler(l) 
	ph := handlers.NewProductsHandler(l)   
	lh := handlers.NewLibraryHandler(l)                                            

//---[ SERVICE PATHS ]---------------------------------------------------------------------------------------------------------------------------------//

	SgetRouther := r.Methods(http.MethodGet).Subrouter()
	SgetRouther.HandleFunc("/students/get/", sh.GetStudents)

	SpostRouther := r.Methods(http.MethodPost).Subrouter()
	SpostRouther.HandleFunc("/students/post/", sh.PostStudents)
	SpostRouther.Use(sh.MiddlewareStudentValidation)

	SputRouther := r.Methods(http.MethodPut).Subrouter()
	SputRouther.HandleFunc("/students/put/{id:[0-9]+}", sh.PutStudents)
	SputRouther.Use(sh.MiddlewareStudentValidation)

	SdeleteRouther := r.Methods(http.MethodDelete).Subrouter()
	SdeleteRouther.HandleFunc("/students/delete/{id:[0-9]+}", sh.DeleteStudents)
	SdeleteRouther.Use(sh.MiddlewareStudentValidation)

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

//---[ SERVER CONFIGRATIONS & LOGGING ]-----------------------------------------------------------------------------------------------------------------//

	cors := goHandlers.CORS(
		goHandlers.AllowedOrigins(config.CORS.AllowedOrigins),
		goHandlers.AllowedMethods(config.CORS.AllowedMethods),
		goHandlers.AllowedHeaders(config.CORS.AllowedHeaders),
	)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler:      cors(r),
		IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	}

	l.Printf("%s[INFO]%s : %sServer Starting at %s:%d%s", cmd.BYELLOW_BLACK, cmd.TRESET, cmd.TVIOLET, config.Server.Host , config.Server.Port, cmd.TRESET)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan

	l.Printf("%s[INFO]%s : %sReceived terminate signal, graceful shutdown %v\n %s", cmd.BYELLOW_BLACK, cmd.TRESET, cmd.TRED, sig, cmd.TRESET)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if err := server.Shutdown(tc); err != nil {
		l.Fatalf("%s[INFO]%s : %serror during server shutdown: %v%s", cmd.BYELLOW_BLACK, cmd.TRESET, cmd.TRED ,err, cmd.TRESET)
	}
	
}