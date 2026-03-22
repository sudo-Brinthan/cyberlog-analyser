package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

func SendToPython(logMsg string, url string) (map[string]interface{}, error) {
	reqBody := Request{
		Message: logMsg,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("❌ Error decoding response:", err)
		return nil, err
	}

	return result, nil
}