package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"context"
	"syscall"
	"net/http"
	"os/signal"
	
	"github.com/AtahanPoyraz/cmd"
	"github.com/AtahanPoyraz/auth"
	"github.com/AtahanPoyraz/router"
	"github.com/AtahanPoyraz/config"
	
	"github.com/gorilla/mux"
	goHandlers "github.com/gorilla/handlers"
)

var (
	r = mux.NewRouter()
	l = log.New(os.Stdout, fmt.Sprintf("%sGOMICRON-API >> %s", cmd.TCYAN, cmd.TRESET), log.LstdFlags) 

	srvh = auth.ServerHandler(l)
)

func main() {
	config, err := config.ReadConfigFromFile("./config.yml")
	
	if err != nil {
		l.Printf("%s[ERROR]%s : Read operation failed please check file : (%e)", cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(1)
	}

	loadStaticFiles(r)
	joinServerURL(r)

	router.SetRoutes(r, l)
	startServer(config, r, l)
}

func loadStaticFiles(r *mux.Router) {
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css/"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js/"))))
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("./scripts/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates/assets/"))))
}

func joinServerURL(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/html/main.html")
	})
	r.HandleFunc("/dashboard/", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/html/dashboard.html")
	})))

	authRouter := r.Methods(http.MethodPost).Subrouter()
	authRouter.HandleFunc("/gomicron/backend/user/auth/", srvh.AuthUser)	
}

func startServer(config config.Config, r *mux.Router, l *log.Logger) {
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