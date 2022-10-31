package imdb

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMovieStruct(t *testing.T) {
	jsonMovie := []byte(`
		{
			"Title": "Corbett and Courtney Before the Kinetograph",
			"Year": "1894",
			"Rated": "Not Rated",
			"Released": "17 Nov 1894",
			"Runtime": "1 min",
			"Genre": "Short, Sport",
			"Director": "William K.L. Dickson, William Heise",
			"Writer": "N/A",
			"Actors": "James J. Corbett, Peter Courtney",
			"Plot": "James J. Corbett and Peter Courtney meet in a boxing exhibition.",
			"Language": "None",
			"Country": "United States",
			"Awards": "N/A",
			"Poster": "https://m.media-amazon.com/images/M/MV5BODA4ZTI3NWYtNzAyYi00Yjc5LTk0NDUtZWIzNzQzYzdkODgyXkEyXkFqcGdeQXVyNTM3MDMyMDQ@._V1_SX300.jpg",
			"Ratings": [
				{
					"Source": "Internet Movie Database",
					"Value": "5.5/10"
				}
			],
			"Metascore": "N/A",
			"imdbRating": "5.5",
			"imdbVotes": "789",
			"imdbID": "tt0000007",
			"Type": "movie",
			"DVD": "N/A",
			"BoxOffice": "N/A",
			"Production": "N/A",
			"Website": "N/A",
			"Response": "True"
		}
	`)
	var m *Movie
	err := json.Unmarshal(jsonMovie, &m)
	if err != nil {
		t.Errorf("unmarshall error %v", err)
	}
	expImbId := "tt0000007"
	if m.ImdbID != expImbId {
		t.Errorf("expected %s imdbid got %s", expImbId, m.ImdbID)
	}
}

func TestMovie_String(t *testing.T) {
	m := Movie{
		Title:  "Corbett and Courtney Before the Kinetograph",
		Plot:   "James J. Corbett and Peter Courtney meet in a boxing exhibition.",
		ImdbID: "tt0000007",
	}
	movieStr := m.String()
	if movieStr != fmt.Sprintf("%s\t %s\t %s\t", m.ImdbID, m.Title, m.Plot) {
		t.Errorf("malformed movie string representation")
	}
}