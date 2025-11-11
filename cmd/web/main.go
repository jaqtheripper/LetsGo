package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// To use environmental variables -> os.Getenv()
	addr := flag.String("addr", "127.0.0.1:4000", "HTTP network address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
