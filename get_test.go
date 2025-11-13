package kokkeneskoekken

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	client := givenMockClient(t, http.StatusOK)
	schedule, err := GetSchedule(client, "::fake::id", testOfferId)
	assert.Nil(t, err)

	assert.Equal(t, "Græsgris", schedule["2025-11-13"]["Dagens ret"][0])
}
