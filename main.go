package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/rkperes/pokitalo/adapters/handlers"
	"github.com/rkperes/pokitalo/adapters/repository"
	"github.com/rkperes/pokitalo/internal/core/services"
)

func main() {
	_ = repository.NewSQLite()
	_ = services.New()
	httpsrv := handlers.NewHTTP()

	addr := ":8080"
	slog.Info("starting server", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, httpsrv); err != nil {
		log.Fatal(err)
	}
}
