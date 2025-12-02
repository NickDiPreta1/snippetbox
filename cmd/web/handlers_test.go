package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NickDiPreta1/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	ping(rr, r)

	res := rr.Result()
	assert.Equal(t, res.StatusCode, http.StatusOK)

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, string(resBytes), "OK")
}
