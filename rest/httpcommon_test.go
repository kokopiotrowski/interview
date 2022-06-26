package rest_test

import (
	"interview/rest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendJSON(t *testing.T) {
	testMessage := "TestMessage"
	recorder := httptest.NewRecorder()
	rest.SendResponseJSON(recorder, http.StatusOK, testMessage)

	assert.Contains(t, recorder.Body.String(), testMessage)
}
