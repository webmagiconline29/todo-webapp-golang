package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddTodo(t *testing.T) {
	req, err := http.NewRequest("POST", "/add", strings.NewReader("todo=TestTodo"))
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

	if todos[0] != "TestTodo" {
		t.Errorf("Expected 'TestTodo', got '%s'", todos[0])
	}
}

func TestTodoList(t *testing.T) {
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

	expected := "<h1>TODO List</h1>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddTodo(t *testing.T) {
	// Reset the global todos variable before each test
	todos = nil

	// Create a new request with a POST method and form value
	req, err := http.NewRequest("POST", "/add", strings.NewReader("todo=TestTodo"))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler function for testing
	handler := http.HandlerFunc(addTodo)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check if the handler returned the correct status code
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check if the todo was added correctly
	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}

	if todos[0] != "TestTodo" {
		t.Errorf("Expected 'TestTodo', got '%s'", todos[0])
	}
}

func TestTodoList(t *testing.T) {
	// Reset the global todos variable before each test
	todos = []string{}

	// Create a new request with a GET method
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler function for testing
	handler := http.HandlerFunc(todoList)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check if the handler returned the correct status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the response body contains the expected HTML
	expected := "<h1>TODO List</h1>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
