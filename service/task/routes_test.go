package task

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/yann-fk-21/todo-app/types"
)


func TestRoutes(t *testing.T) {
	taskStore := &mocktaskStore{}
	handler := NewHandler(taskStore)

	t.Run("should return http status 400", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/tasks", nil)
		if err != nil {
			log.Fatal(err)
		}
        rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", handler.createTask).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got %v", http.StatusBadRequest, rr.Code)
		}

	})

	t.Run("should return http status 201", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/tasks", nil)
		if err != nil {
			log.Fatal(err)
		}
        rr := httptest.NewRecorder()
		router := mux.NewRouter()

		task := types.Task{
			Title: "Hello",
			Description: "heeeeeee",
			Status: true,
			CreatedAt: time.Now(),
		}

		taskJSON, err := json.Marshal(task)
		if err != nil {
			log.Fatal(err)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(taskJSON))

		router.HandleFunc("/tasks", handler.createTask).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %v, got %v", http.StatusCreated, rr.Code)
		}

	})
}

type mocktaskStore struct {}

func (m *mocktaskStore) CreateTask(t types.Task) error {
	return nil
}

func (m *mocktaskStore) GetTasks()([]types.Task, error) {
	return nil, nil
}