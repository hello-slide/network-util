package networkutil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetHeader(w http.ResponseWriter, req *http.Request) (map[string]string, error) {
	if req.Method != "POST" {
		ErrorStatus(w)
		return nil, fmt.Errorf("bad request")
	}

	if req.Header.Get("Content-Type") != "application/json" {
		ErrorStatus(w)
		return nil, fmt.Errorf("bad request")
	}

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		ErrorStatus(w)
		return nil, fmt.Errorf("content-length flags is not found")
	}

	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		ErrorStatus(w)
		return nil, fmt.Errorf("parse error")
	}

	var jsonBody map[string]string
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		ErrorStatus(w)
		fmt.Println(err)
		return nil, fmt.Errorf("parse error")
	}
	return jsonBody, nil
}

func PickValue(key string, data map[string]string, w http.ResponseWriter) (string, error) {
	if value, ok := data[key]; ok {
		return value, nil
	}
	ErrorStatus(w)
	return "", fmt.Errorf("Key is not found")
}

func GetFromKey(key string, w http.ResponseWriter, req *http.Request) (string, error) {
	body, err := GetHeader(w, req)
	if err != nil {
		return "", err
	}
	return PickValue(key, body, w)
}
