package main

import (
	"encoding/json"
	"fmt"
	"ginblog/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggerFileWrite(t *testing.T) {
	token, err := GetToken(TestUser.username, TestUser.password)
	if err != nil {
		fmt.Println(err)
	}
	router := routers.InitRouter()
	uri := "/api/v1/articles?token=" + token + "&tag_id=0"
	responseData := make(map[string]interface{})

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	router.ServeHTTP(writer, req)

	decodeErr := json.NewDecoder(writer.Body).Decode(&responseData)
	if err != nil {
		fmt.Println(decodeErr)
	}

	code := int(responseData["code"].(float64))

	assert.Equal(t, http.StatusBadRequest, code)
}
