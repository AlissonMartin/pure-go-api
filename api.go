package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

type APIServer struct {
	listenAddr string
	store Storage
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

func MakeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		} 
	}
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	
	router.HandleFunc("/account", MakeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", MakeHTTPHandleFunc(s.handleAccount))

	fmt.Println("Server running at " + s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w  http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w  http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleCreateAccount(w  http.ResponseWriter, r *http.Request) error {
	CreateAccountReq := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(CreateAccountReq); err != nil {
		return err
	}
	
	account := NewAccount(CreateAccountReq.Name)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	} 
	
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) GetAccounts(w http.ResponseWriter, r *http.Request) error {
	return nil
}
