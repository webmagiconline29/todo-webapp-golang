// main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// func TestAddTodo(t *testing.T) {
// 	resetTodos()

// 	req, err := http.NewRequest("POST", "/add", strings.NewReader("todo=TestTodo"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(addTodo)

// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusSeeOther {
// 		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
// 	}

// 	if len(todos) != 1 {
// 		t.Errorf("Expected 1 todo, got %d", len(todos))
// 	}

// 	if todos[0] != "TestTodo" {
// 		t.Errorf("Expected 'TestTodo', got '%s'", todos[0])
// 	}
// }

func TestTodoList(t *testing.T) {
	resetTodos()

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

	expected := "<h1>WebMagic TODO List</h1>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func resetTodos() {
	todos = []string{}
}
