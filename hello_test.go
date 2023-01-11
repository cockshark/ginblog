package main

import (
	"encoding/binary"
	"ginblog/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HelloWorld struct {
	message string
}

func TestHelloWord(t *testing.T) {

	var helloWorld HelloWorld

	router := routers.InitRouter()

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(writer, req)

	err := binary.Read(writer.Body, binary.BigEndian, &helloWorld)
	if err != nil {
		return
	}
	if err != nil {
		t.Log(err)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "test", helloWorld.message)
}
