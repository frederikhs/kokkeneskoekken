package kokkeneskoekken

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"testing"
)

type mockTransport struct {
	RoundTripFunc func(req *http.Request) *http.Response
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req), nil
}

func givenMockClient(t *testing.T, statusCode int) *http.Client {
	data, err := os.ReadFile("testfiles/data.min.json")
	assert.Nil(t, err)

	mockResp := &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewReader(data)),
		Header:     make(http.Header),
	}

	client := &http.Client{
		Transport: &mockTransport{
			RoundTripFunc: func(req *http.Request) *http.Response {
				return mockResp
			},
		},
	}

	return client
}

func TestFetch(t *testing.T) {
	client := givenMockClient(t, http.StatusOK)
	responseBody, err := fetchForSchoolIdWithClient(client, "::fake::id")
	assert.Nil(t, err)

	schedule, err := parse(responseBody, testOfferId)
	assert.Nil(t, err)

	assert.Equal(t, "Græsgris", schedule["2025-11-13"]["Dagens ret"][0])
}

func TestFetchBadStatusCode(t *testing.T) {
	client := givenMockClient(t, http.StatusInternalServerError)
	responseBody, err := fetchForSchoolIdWithClient(client, "::fake::id")
	assert.NotNil(t, err)
	assert.Nil(t, responseBody)
}
