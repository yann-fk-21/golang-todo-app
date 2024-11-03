package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yann-fk-21/todo-app/logger"
	"github.com/yann-fk-21/todo-app/service/task"
)

type Server struct {
	Addr string
	db *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		Addr: addr,
		db: db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	taskStore := task.NewStore(s.db, logger.InitLogger())

	taskHandler := task.NewHandler(taskStore)
	
	taskHandler.RegisterHandlerRoutes(subRouter)

	fmt.Printf("Server run on port %s", s.Addr)
	return http.ListenAndServe(s.Addr, router)
}