package models

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is the way of defining unit tests in go.
func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id=0")
	// The above line is doing the same thing as below.
	// if user != nil {
	// 	t.Error("we were not expecting a user with id=0")
	// }
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	// if err == nil {
	// 	t.Error("we were expecting  an error  when user id=0")
	// }
}

func TestGetUserNoError(t *testing.T) {

	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Pralay", user.FirstName)
	assert.EqualValues(t, "Saw", user.LastName)
	assert.EqualValues(t, "myemail@hotmail.com", user.Email)
}
