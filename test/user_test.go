package test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"main/domain"
	"main/handler"
	"net/http"
	"net/http/httptest"
	testing "testing"
)

var userPayload = `{"id": 1, "name": "test", "password": "test"}`

func TestGetUser(t *testing.T) {
	// 	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assume you have a function GetUser that handles the request
	if assert.NoError(t, handler.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		user := new(domain.User)
		err := json.Unmarshal(rec.Body.Bytes(), user)
		assert.NoError(t, err)
	}
}

func TestCreateUser(t *testing.T) {
	// setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/user",
		bytes.NewReader([]byte(userPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assume you have a function CreateUser that handles the request
	if assert.NoError(t, handler.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		user := new(domain.User)
		err := json.Unmarshal(rec.Body.Bytes(), user)
		assert.NoError(t, err)

		// Verify the user details here
		assert.Equal(t, "test", user.Name)
		//... other assertions
	}
}
