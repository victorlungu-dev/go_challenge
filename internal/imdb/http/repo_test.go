package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockResponse(t *testing.T) []byte {
	t.Helper()
	return []byte(`{
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
		}`)
}

func invalidMockResp(t *testing.T) []byte {
	return []byte(`<some invalid data>`)
}

func TestHttpRepo_Retrieve(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("apiKey") == "" {
			t.Errorf("expected apiKey param in url query")
		}
		if q.Get("i") == "" {
			t.Errorf("missing mandatory parameter i")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponse(t))
	}))
	defer server.Close()
	repo := NewHttpRepo(context.Background(), server.URL, "testApiKey", 1)
	expId := "tt0000007"
	m, err := repo.Retrieve(expId)
	if err != nil {
		t.Errorf("%v", err)
	}
	if m.ImdbID != expId {
		t.Errorf("Expected id %s got %s", expId, m.ImdbID)
	}
	_, err = repo.Retrieve(expId)
	if err == nil {
		t.Errorf("maxRequests error should have been given")
	}
}

func TestHttpRepo_RetrieveInvalidData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("apiKey") == "" {
			t.Errorf("expected apiKey param in url query")
		}
		if q.Get("i") == "" {
			t.Errorf("missing mandatory parameter i")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(invalidMockResp(t))
	}))
	defer server.Close()
	repo := NewHttpRepo(context.Background(), server.URL, "testApiKey", 1)
	expId := "tt0000007"

	_, err := repo.Retrieve(expId)
	if err == nil {
		t.Errorf("should have given an error on unmarshall the body")
	}

}
