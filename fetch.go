package kokkeneskoekken

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const url = "https://kokkeneskoekken.kanpla.dk/api/internal/load/frontend"

type requestBody struct {
	SchoolId string `json:"schoolId"`
}

func fetchForSchoolIdWithClient(client *http.Client, schoolId string) ([]byte, error) {
	rb := requestBody{SchoolId: schoolId}
	requestBody, err := json.Marshal(&rb)
	if err != nil {
		panic(err)
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("did not get status ok, got: " + resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
