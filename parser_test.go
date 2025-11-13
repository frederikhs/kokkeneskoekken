package kokkeneskoekken

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const testOfferId = "abc123"

func givenScheduleByFilename(t *testing.T, filename string) Schedule {
	data, err := os.ReadFile(filename)
	assert.Nil(t, err)

	schedule, err := parse(data, testOfferId)
	assert.Nil(t, err)

	return schedule
}

func TestParse(t *testing.T) {
	schedule := givenScheduleByFilename(t, "testfiles/data.min.json")

	assert.Len(t, schedule["2025-11-13"]["Dagens ret"], 1)
	assert.Equal(t, "Græsgris", schedule["2025-11-13"]["Dagens ret"][0])
	assert.Len(t, schedule["2025-11-13"]["Dagens vegetariske ret"], 1)
	assert.Equal(t, "Spidskål", schedule["2025-11-13"]["Dagens vegetariske ret"][0])
}

func TestParseMultipleOfSameDishType(t *testing.T) {
	schedule := givenScheduleByFilename(t, "testfiles/data.min2.json")

	assert.Len(t, schedule["2025-11-13"]["Pålæg"], 2)
	assert.Equal(t, "Leverpostej", schedule["2025-11-13"]["Pålæg"][0])
	assert.Equal(t, "Æg og rejer", schedule["2025-11-13"]["Pålæg"][1])
}

func TestParseMultipleDaysInSchedule(t *testing.T) {
	schedule := givenScheduleByFilename(t, "testfiles/data.min3.json")

	assert.Len(t, schedule["2025-11-13"]["Pålæg"], 1)
	assert.Len(t, schedule["2025-11-14"]["Pålæg"], 1)
	assert.Equal(t, "Leverpostej", schedule["2025-11-13"]["Pålæg"][0])
	assert.Equal(t, "Æg og rejer", schedule["2025-11-14"]["Pålæg"][0])
}

func TestEmptyMenuName(t *testing.T) {
	schedule := givenScheduleByFilename(t, "testfiles/data.min4.json")

	assert.Len(t, schedule["2025-11-13"], 0)
}

func TestNotAvailable(t *testing.T) {
	schedule := givenScheduleByFilename(t, "testfiles/data.min5.json")

	assert.Len(t, schedule["2025-11-13"], 0)
}

func TestBadStructure(t *testing.T) {
	data, err := os.ReadFile("testfiles/data.min6.json")
	assert.Nil(t, err)

	schedule, err := parse(data, testOfferId)
	assert.Nil(t, err)
	assert.Len(t, schedule, 0)
}
