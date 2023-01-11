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

func getAllTags() (code int, tags map[string]interface{}) {
	token, err := GetToken(TestUser.username, TestUser.password)
	if err != nil {
		fmt.Println(err)
	}
	router := routers.InitRouter()
	uri := "/api/v1/tags?token=" + token
	responseData := make(map[string]interface{})

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	router.ServeHTTP(writer, req)

	decodeErr := json.NewDecoder(writer.Body).Decode(&responseData)
	if err != nil {
		fmt.Println(decodeErr)
	}

	fmt.Println(responseData)
	code = int(responseData["code"].(float64))
	tags = responseData["data"].(map[string]interface{})

	return

}

func TestApiGetTags(t *testing.T) {
	code, _ := getAllTags()
	assert.Equal(t, http.StatusOK, code, "%/api/v1/tags response status is : %s", code)

}

func TestTagOption(t *testing.T) {
	// get all tags
	// filter target tag id
	// delete tag
	// add tag
	// add tag

	token, err := GetToken(TestUser.username, TestUser.password)
	if err != nil {
		t.Error(err)
	}
	router := routers.InitRouter()
	uri := "/api/v1/tags?token=" + token

	name := "wushuhua"
	createdBy := "wushuhua"

	// add tag
	addUrl := uri + "&name=" + name + "&state=1" + "&created_by=" + createdBy
	addWriter := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", addUrl, nil)
	router.ServeHTTP(addWriter, req)

	addResponseData := make(map[string]interface{})
	AddDecodeErr := json.NewDecoder(addWriter.Body).Decode(&addResponseData)
	if AddDecodeErr != nil {
		t.Error(AddDecodeErr)
	}
	addTagRespCode := int(addResponseData["code"].(float64))

	assert.Equal(t, http.StatusOK, addTagRespCode, "%s response status is : %s", uri, addTagRespCode)

}
