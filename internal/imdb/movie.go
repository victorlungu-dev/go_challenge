package imdb

import "fmt"

type Movie struct {
	Title      string   `json:"Title,omitempty"`
	Year       string   `json:"Year,omitempty"`
	Rated      string   `json:"Rated,omitempty"`
	Released   string   `json:"Released,omitempty"`
	Runtime    string   `json:"Runtime,omitempty"`
	Genre      string   `json:"Genre,omitempty"`
	Director   string   `json:"Director,omitempty"`
	Writer     string   `json:"Writer,omitempty"`
	Actors     string   `json:"Actors,omitempty"`
	Plot       string   `json:"Plot,omitempty"`
	Language   string   `json:"Language,omitempty"`
	Country    string   `json:"Country,omitempty"`
	Awards     string   `json:"Awards,omitempty"`
	Poster     string   `json:"Poster,omitempty"`
	Ratings    []Rating `json:"Ratings,omitempty"`
	Metascore  string   `json:"Metascore,omitempty"`
	ImdbRating string   `json:"imdbRating,omitempty"`
	ImdbVotes  string   `json:"imdbVotes,omitempty"`
	ImdbID     string   `json:"imdbID,omitempty"`
	Type       string   `json:"Type,omitempty"`
	DVD        string   `json:"DVD,omitempty"`
	BoxOffice  string   `json:"BoxOffice,omitempty"`
	Production string   `json:"Production,omitempty"`
	Website    string   `json:"Website,omitempty"`
	Response   string   `json:"Response,omitempty"`
}

type Rating struct {
	Source string `json:"Source,omitempty"`
	Value  string `json:"Value,omitempty"`
}

func (m *Movie) String() string {
	return fmt.Sprintf("%s\t %s\t %s\t", m.ImdbID, m.Title, m.Plot)
}
