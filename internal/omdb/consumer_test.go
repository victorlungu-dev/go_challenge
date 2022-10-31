package omdb

import (
	"azarc.io/internal/argparser"
	"azarc.io/internal/imdb"
	"errors"
	"regexp"
	"testing"
)

type MockRepo struct{}

func (mr *MockRepo) Retrieve(id string) (*imdb.Movie, error) {
	m := imdb.Movie{
		Title:  "Corbett and Courtney Before the Kinetograph",
		Plot:   "James J. Corbett and Peter Courtney meet in a boxing exhibition.",
		ImdbID: "tt0000007",
	}
	if id != "tt0000007" {
		return nil, errors.New("invalid id")
	}
	return &m, nil
}
func TestConsumerWorker(t *testing.T) {
	inChan := make(chan string, 1)
	outChan := make(chan *imdb.Movie, 1)
	quit := make(chan bool, 1)
	numGoroutines := 1
	re, _ := regexp.Compile("")
	config := argparser.Configuration{
		Opts:         argparser.Options{},
		Filters:      argparser.StringFilters{},
		RegExpFilter: re,
		ApiKey:       "",
	}
	consumer := Consumer{
		InChannel:  inChan,
		OutChannel: outChan,
		Quit:       quit,
		Workers:    numGoroutines,
		Repo:       &MockRepo{},
		Header:     Header,
		Config:     &config,
	}
	go consumer.Consume()

	inChan <- "tt0000007\ttitle\torigtitle\tplot"
	m := <-outChan
	expId := "tt0000007"
	if m.ImdbID != expId {
		t.Errorf("expected id %s got %s", expId, m.ImdbID)
	}

	inChan <- "tt0000008\ttitle\torigtitle\tplot"
	closing := <-quit
	if !closing {
		t.Errorf("expected closing signal")
	}

	quit <- true
	<-quit

}
