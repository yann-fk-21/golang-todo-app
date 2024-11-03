package task

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yann-fk-21/todo-app/types"
	"github.com/yann-fk-21/todo-app/utils"
)

type Handler struct {
	store types.TaskStore
}

func NewHandler(store types.TaskStore) *Handler {
 return &Handler{
	store: store,
 }
}

func (h *Handler) RegisterHandlerRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", h.createTask).Methods("POST")
	router.HandleFunc("/tasks", h.getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", h.getTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", h.updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", h.deleteTask).Methods("DELETE")
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
 
var task types.Task
err := utils.ParseJSON(r, &task)

if err != nil {
	utils.WriteError(w, http.StatusBadRequest, err)
	return
}

err = h.store.CreateTask(task)
if err != nil {
	utils.WriteError(w, http.StatusInternalServerError, err)
	return
}

utils.WriteJSON(w, http.StatusCreated, "task created successfully")

}

func (h *Handler) getTasks(w http.ResponseWriter, r *http.Request) {
// TODO
}

func (h *Handler) getTaskById(w http.ResponseWriter, r *http.Request) {
// TODO
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
// TODO
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
// TODO
}