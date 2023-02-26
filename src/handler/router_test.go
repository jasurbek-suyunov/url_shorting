package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func decodeJSON[T any](body io.Reader) T {
	b, _ := io.ReadAll(body)

	var s T
	json.Unmarshal(b, &s)

	return s
}

func TestLogin(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()

	body := strings.NewReader("{ \"email\": \"foo\", \"password\": \"bar\" }")

	req, _ := http.NewRequest("POST", "/login", body)
	r.ServeHTTP(w, req)

	result := decodeJSON[struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}](w.Body)

	fmt.Println(result)
}
