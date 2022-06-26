package rest

import (
	"bytes"
	"encoding/json"
	"interview/conf"
	"io"
	"net/http"
	"time"
)

type Message struct {
	Message string `json:"message"`
}

type ResponseToClient struct {
	Echo        string `json:"echo"`
	Timestamp   int64  `json:"timestamp"`
	Environment string `json:"env"`
	Version     string `json:"version"`
}

func PingPostman(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "",
	}
	DecodeBodyAsJSON(r, &message)

	responseFromPostman, err := EchoPostman(message.Message)
	if err != nil {
		SendResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := buildPingPostmanResponse(responseFromPostman)

	SendResponseJSON(w, http.StatusOK, response)
}

func EchoPostman(message string) (string, error) {
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(message)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	requestToPostman, err := http.NewRequest(http.MethodPost, "https://postman-echo.com/post", &buffer)
	if err != nil {
		return "", err
	}

	response, err := client.Do(requestToPostman)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyFromPostmanBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bodyFromPostmanBytes), nil
}

func buildPingPostmanResponse(responseFromPostman string) ResponseToClient {
	conf := conf.GetConf()
	currentTimeInUnix := time.Now().Unix()

	responseToClient := ResponseToClient{
		Echo:        responseFromPostman,
		Timestamp:   currentTimeInUnix,
		Version:     conf.Version,
		Environment: conf.Environment,
	}
	return responseToClient
}
