package kokkeneskoekken

import (
	"encoding/json"
	"strings"
	"time"
)

func parse(responseBody []byte, offerId string) (Schedule, error) {
	schedule := make(Schedule)
	var response apiResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	for _, b := range response.Offers[offerId].Items {
		for t, c := range b.Dates {
			if !c.Available {
				continue
			}

			if len(c.Menu.Name) == 0 {
				continue
			}

			aDate := time.Unix(t, 0).Format("2006-01-02")

			_, exist := schedule[aDate]
			if !exist {
				schedule[aDate] = make(Menu)
			}

			dishGroup := strings.TrimSpace(b.Name)

			_, exist = schedule[aDate][dishGroup]
			if !exist {
				schedule[aDate][dishGroup] = []string{c.Menu.Name}
			} else {
				schedule[aDate][dishGroup] = append(schedule[aDate][dishGroup], c.Menu.Name)
			}
		}
	}

	return schedule, nil
}
