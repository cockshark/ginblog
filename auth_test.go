package main

import (
	"encoding/json"
	"errors"
	"ginblog/pkg/e"
	"ginblog/pkg/util"
	"ginblog/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var TestUser User = User{
	username: "test",
	password: "test123456",
}

type User struct {
	username string
	password string
}

func GetToken(username, password string) (string, error) {
	router := routers.InitRouter()
	var uri string = "/auth?username=" + username + "&password=" + password
	tokenMap := make(map[string]string)
	responseData := make(map[string]interface{})

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	router.ServeHTTP(writer, req)

	err := json.NewDecoder(writer.Body).Decode(&responseData)
	if err != nil {
		return "", err
	}

	for _, val := range responseData {
		switch val.(type) {
		case map[string]interface{}:
			for _, token := range val.(map[string]interface{}) {
				tokenMap["token"] = token.(string)
			}
		}
	}

	if writer.Code == http.StatusOK {
		return tokenMap["token"], nil
	}

	return "", errors.New(e.MsgFlags[e.INVALID_PARAMS])

}

func TestAuthToken(t *testing.T) {

	token, err := GetToken(TestUser.username, TestUser.password)
	assert.Equal(t, nil, err, "fail to get token :%s", err)

	claims, err := util.ParseToken(token)
	if err != nil {
		t.Log(err)
	}

	assert.Equal(t, TestUser.username, claims.Username, "username unequal")
	assert.Equal(t, TestUser.password, claims.Password, "password unequal")

}
