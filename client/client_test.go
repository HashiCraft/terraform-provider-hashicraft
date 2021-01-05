package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func setupTestServer(t *testing.T, key, contentType, path, body *string, resp string) string {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*path = r.URL.Path
		*key = r.Header.Get("authorization")
		*contentType = r.Header.Get("content-type")

		d, _ := ioutil.ReadAll(r.Body)
		*body = string(d)

    fmt.Fprint(w, resp)
	}))

	t.Cleanup(func() {
		ts.Close()
	})

	return ts.URL
}

func TestHealthEndpoint(t *testing.T) {
	var key string
	var path string
	var content string
	var body string

	url := setupTestServer(t, &key, &content, &path, &body, "OK")
	c := New(url, "abc")

	resp, err := c.Health()

	assert.NoError(t, err)
	assert.Equal(t, "OK", resp)
	assert.Equal(t, "/health", path)
	assert.Equal(t, "bearer abc", key)
}

func TestNewEndpoint(t *testing.T) {
	var key string
	var path string
	var content string
	var body string

  url := setupTestServer(t, &key, &content, &path, &body, `{"id": "sdf"}`)
	c := New(url, "abc")

	conf := NewRequest{
    Host: "abc",
    Port:1234,
    Username:"123",
    Password:"456",
	}

	resp, err := c.New(conf)

	assert.NoError(t, err)
	assert.Equal(t, "/bot", path)
	assert.Equal(t, "bearer abc", key)

  assert.JSONEq(t, `{
    "host": "abc",
    "port": 1234,
    "username": "123",
    "password": "456" 
  }`,  body)

  assert.Equal(t, "sdf", resp)
}

func TestDeleteEndpoint(t *testing.T) {
	var key string
	var path string
	var content string
	var body string

	url := setupTestServer(t, &key, &content, &path, &body, "")
	c := New(url, "abc")

	err := c.Delete("abc")

	assert.NoError(t, err)
	assert.Equal(t, "/bot/abc", path)
}

func TestConfigureEndpoint(t *testing.T) {
	var key string
	var path string
	var content string
	var body string

	url := setupTestServer(t, &key, &content, &path, &body, "")
	c := New(url, "abc")

	conf := ConfigRequest{
		MineStart: "123,1,123",
		MineEnd:   "123,1,123",
		ToolChest: "123,1,123",
		DropChest: "123,1,123",
	}

	err := c.Configure("abc",conf)

	assert.NoError(t, err)
	assert.Equal(t, "/bot/abc/configure", path)
	assert.Equal(t, "bearer abc", key)

  assert.JSONEq(t, `{
    "mine_start": "123,1,123",
    "mine_end": "123,1,123",
    "tool_chest": "123,1,123",
    "drop_chest": "123,1,123"
  }`,  body)
}
