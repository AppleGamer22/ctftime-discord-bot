package api

import (
	"time"
)

type Event struct {
	Description       string    `json:"description"`
	Format            string    `json:"format"`
	CTFTimeURL        string    `json:"ctftime_url"`
	LiveFeed          string    `json:"live_feed"`
	Logo              string    `json:"logo"`
	URL               string    `json:"url"`
	UniversityWebsite string    `json:"university_website"`
	University        string    `json:"university"`
	Restrictions      string    `json:"restrictions"`
	Participants      uint      `json:"participants"`
	Weight            float32   `json:"weight"`
	Start             time.Time `json:"start"`
	Finish            time.Time `json:"finish"`
	Durration         struct {
		Hours uint `json:"hours"`
		Days  uint `json:"days"`
	} `json:"duration"`
	ID    uint `json:"id"`
	CTFID uint `json:"ctf_id"`
}

type Team struct {
	Academic     bool     `json:"academic"`
	ID           uint     `json:"id"`
	PrimaryAlias string   `json:"primary_alias"`
	Name         string   `json:"name"`
	Logo         string   `json:"logo"`
	Country      string   `json:"country"`
	Aliases      []string `json:"aliases"`
	Rating       map[string]struct {
		RatingPlace     uint    `json:"rating_place"`
		OrganizerPoints float32 `json:"organizer_points"`
		RatingPoints    float32 `json:"rating_points"`
		CountryPlace    uint    `json:"country_place"`
	} `json:"rating"`
}
