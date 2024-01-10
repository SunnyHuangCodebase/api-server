package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootShouldSucceed(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handleRoot))
	defer s.Close()

	c := s.Client()

	r, err := c.Get(s.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Welcome", string(body))
	assert.Equal(t, http.StatusOK, r.StatusCode)
}

func TestRootShouldFail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handleRoot))
	defer s.Close()

	c := s.Client()

	body := strings.NewReader("")
	r, err := c.Post(s.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusMethodNotAllowed, r.StatusCode)
}

func TestHelloWorldShouldSucceed(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer s.Close()

	c := s.Client()

	r, err := c.Get(s.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello World", string(body))
	assert.Equal(t, http.StatusOK, r.StatusCode)
}

func TestHelloWorldShouldFail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer s.Close()

	c := s.Client()

	body := strings.NewReader("")
	r, err := c.Post(s.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusMethodNotAllowed, r.StatusCode)
}

func TestHealthShouldSucceed(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer s.Close()

	c := s.Client()

	r, err := c.Get(s.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "OK", string(body))
	assert.Equal(t, http.StatusOK, r.StatusCode)
}

func TestHealthShouldFail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer s.Close()

	c := s.Client()

	body := strings.NewReader("")
	r, err := c.Post(s.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusMethodNotAllowed, r.StatusCode)
}
