package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorvk/rent-app/api-services/initializers"
	_ "github.com/gorvk/rent-app/api-services/routes"
)

func main() {
	var err error

	// uncomment below code for local development without docker
	// err = initializers.LoadEnv()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// connecting to DB
	err = initializers.ConnectDB()
	if err != nil {
		errorMsg := err.Error()
		fmt.Println(errorMsg)
		return
	}

	// running migration if DB_MIGRATION_FLAG is Y
	dbMigrationFlag := os.Getenv("DB_MIGRATION_FLAG")
	if dbMigrationFlag == "Y" {
		err = initializers.DbMigration()
		if err != nil {
			errorMsg := err.Error()
			fmt.Println(errorMsg)
		}
		return
	}

	configureListenAndServe()
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
}

// adding cors headers
func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

// configuring server and listening at the given port
func configureListenAndServe() {
	fmt.Println("Starting Server...")

	port := os.Getenv("API_PORT")

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      addCorsHeaders(http.DefaultServeMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		fmt.Println("Server Listening at Port " + port)
		err := server.ListenAndServe()

		if err != nil {
			fmt.Println(err)
		}
	}()

	osSignalChan := make(chan os.Signal)
	signal.Notify(osSignalChan, os.Interrupt, os.Kill)
	osSignal := <-osSignalChan
	fmt.Println("gracefully shutting down server due to :", osSignal)
	contextWithTimeout, cancelContext := context.WithTimeout(context.Background(), 30*time.Second)
	cancelContext()
	server.Shutdown(contextWithTimeout)
}
