// package main

// import (
// 	"fmt"
// 	"net/http"

// 	// "golang.org/x/exp/slog"

// 	"./routes"
// 	"appsec-notifayer/internal/config"
// )

// func main() {
// 	cfg := config.MustLoad()

// 	fmt.Println(cfg)

// 	Test()

// 	router := SetupRoutes()

// 	srv := &http.Server{
// 		Addr:    cfg.Address,
// 		Handler: router,
// 	}
// 	srv.ListenAndServe()
// }
// main.go
package main

import (
	"fmt"
	"net/http"

	"./" // импортируем пакет, содержащий функцию SetupRoutes
	"appsec-notifayer/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	router := SetupRoutes()
	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	srv.ListenAndServe()
}
