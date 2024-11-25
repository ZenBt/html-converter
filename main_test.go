package main

import (
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
)



func TestRequestHandler(t *testing.T) {
  expected := "OK"
  req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
  w := httptest.NewRecorder()
  HealtCheck(w, req)
  res := w.Result()
  defer res.Body.Close()

  data, err := io.ReadAll(res.Body)

  if err != nil {
    t.Errorf("Error: %v", err)
  }
  if res.StatusCode != http.StatusOK {
	t.Errorf("Expected status 200 but got %d", res.StatusCode)
  }
  if string(data) != expected {
    t.Errorf("Expected OK but got %v", string(data))
  }

}
