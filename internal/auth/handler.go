package auth

import (
	"encoding/json"
	"fmt"
	"go-api/configs"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{Config: deps.Config}
	prefix := "/auth"
	router.HandleFunc("POST "+prefix+"/login", handler.login)
	router.HandleFunc("POST "+prefix+"/register", handler.register)
}

func (handler *AuthHandler) login(w http.ResponseWriter, req *http.Request) {
	fmt.Println("login", handler.Config.Auth.Secret)
	res := LoginResponse{
		Token: handler.Config.Auth.Secret,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

func (handler *AuthHandler) register(w http.ResponseWriter, req *http.Request) {
	fmt.Println("register")
}
