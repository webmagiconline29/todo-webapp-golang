package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"sync"
)

func TestAddTodo(t *testing.T) {
	var mu sync.Mutex // Create a Mutex

	mu.Lock()
	defer mu.Unlock()

	resetTodos()

	// ... Rest of your test code ...
}

func TestTodoList(t *testing.T) {
	var mu sync.Mutex // Create a Mutex

	mu.Lock()
	defer mu.Unlock()

	resetTodos()

func TestTodoList(t *testing.T) {
	// Reset todos before testing
	todos = []string{"Task 1", "Task 2"}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoList)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Task 1"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	expected = "Task 2"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestAddTodo(t *testing.T) {
	resetTodos()

	formData := "todo=New Task"

	req, err := http.NewRequest("POST", "/add", strings.NewReader(formData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}

	if todos[0] != "New Task" {
		t.Errorf("Expected 'New Task', got '%s'", todos[0])
	}
}

func resetTodos() {
	mu.Lock()
	defer mu.Unlock()
	todos = []string{}
}
