package main

import (
	"fmt"
	"go-api/configs"
	"go-api/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	addr := conf.Http.Host + ":" + conf.Http.Port

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Println("Server start:", addr)
	server.ListenAndServe()
}
