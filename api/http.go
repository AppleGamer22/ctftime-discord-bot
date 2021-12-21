package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GETRequest(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	request.Header.Set("User-Agent", "MonSec")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func UpcomingEvents(limit *string) ([]Event, error) {
	url := func() string {
		if limit != nil {
			return fmt.Sprintf("https://ctftime.org/api/v1/events/?limit=%s", *limit)
		} else {
			return "https://ctftime.org/api/v1/events/"
		}
	}()
	body, err := GETRequest(url)
	if err != nil {
		return []Event{}, err
	}
	var events []Event
	err = json.Unmarshal(body, &events)
	return events, err
}

func TeamInfo(teamID string) (Team, error) {
	url := fmt.Sprintf("https://ctftime.org/api/v1/teams/%s/", teamID)
	body, err := GETRequest(url)
	if err != nil {
		return Team{}, err
	}
	var team Team
	err = json.Unmarshal(body, &team)
	return team, err
}
