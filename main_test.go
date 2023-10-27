package main

import (
  "net/http"
  "testing"
)

func TestTodoList(t *testing.T) {
  // Create a mock HTTP request.
  req, err := http.NewRequest("GET", "/", nil)
  if err != nil {
    t.Fatal(err)
  }

  // Create a mock HTTP response writer.
  w := http.Recorder{}

  // Call the `todoList` function with the mock request and response writer.
  todoList(w, req)

  // Check that the HTTP status code is 200 OK.
  if w.Code != http.StatusOK {
    t.Errorf("Expected HTTP status code 200 OK, but got %d", w.Code)
  }
}

func TestAddTodo(t *testing.T) {
  // Create a mock HTTP request.
  req, err := http.NewRequest("POST", "/", nil)
  if err != nil {
    t.Fatal(err)
  }

  // Set the form value for the "todo" field.
  req.Form.Set("todo", "Test todo")

  // Create a mock HTTP response writer.
  w := http.Recorder{}

  // Call the `addTodo` function with the mock request and response writer.
  addTodo(w, req)

  // Check that the HTTP status code is 303 See Other.
  if w.Code != http.StatusSeeOther {
    t.Errorf("Expected HTTP status code 303 See Other, but got %d", w.Code)
  }

  // Check that the new todo item was added to the `todos` slice.
  if len(todos) != 1 {
    t.Errorf("Expected 1 todo item in the `todos` slice, but got %d", len(todos))
  }

  // Check that the new todo item has the value "Test todo".
  if todos[0] != "Test todo" {
    t.Errorf("Expected the first todo item to be \"Test todo\", but got %s", todos[0])
  }
}