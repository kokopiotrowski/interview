package rest_test

import (
	"interview/rest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostmanConnection(t *testing.T) {
	testMessage := "TestMessage"
	response, err := rest.EchoPostman(testMessage)
	if assert.NoError(t, err) {
		if !strings.Contains(response, testMessage) {
			t.FailNow()
		}
	}
}
