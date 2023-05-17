package jsonRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostJsonReq(requestRoute string, requestObject interface{}) (*http.Request, error) {
	req, err := http.NewRequest("POST", requestRoute, nil) // "/"
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	json, err := json.Marshal(requestObject)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(json))

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
