package kokkeneskoekken

import (
	"net/http"
)

func GetSchedule(client *http.Client, schoolId string, offerId string) (Schedule, error) {
	responseBody, err := fetchForSchoolIdWithClient(client, schoolId)
	if err != nil {
		return nil, err
	}

	return parse(responseBody, offerId)
}
