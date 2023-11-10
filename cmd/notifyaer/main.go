package main

import (
	"fmt"
	"net/http"

	// "golang.org/x/exp/slog"

	"appsec-notifayer/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	Test()

	router := SetupRoutes()

	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	srv.ListenAndServe()
}
