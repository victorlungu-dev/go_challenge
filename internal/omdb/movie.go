package omdb

import (
	"strings"
)

type Movie struct {
	Tconst         string
	TitleType      string
	PrimaryTitle   string
	OriginalTitle  string
	IsAdult        string
	StartYear      string
	EndYear        string
	RuntimeMinutes string
	Genres         []string
}

// NewMovieFromString build a new Movie struct from a formatted string
// s should be a valid formatted string
func NewMovieFromString(s string) *Movie {
	parts := strings.Split(s, "\t")
	return &Movie{
		Tconst:         parts[0],
		TitleType:      parts[1],
		PrimaryTitle:   parts[2],
		OriginalTitle:  parts[3],
		IsAdult:        parts[4],
		StartYear:      parts[5],
		EndYear:        parts[6],
		RuntimeMinutes: parts[7],
		Genres:         strings.Split(parts[8], " "),
	}
}
