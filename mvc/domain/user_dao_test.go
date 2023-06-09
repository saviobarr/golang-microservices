package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {

	user, err := GetUser(0)

	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an error when user id is 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when user is not found")

	}

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "Expecting an error")

}
