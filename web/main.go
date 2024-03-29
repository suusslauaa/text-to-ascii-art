package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Input  string
	Output string
}

type Error struct {
	Message string
	Code    int
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на http://localhost%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
