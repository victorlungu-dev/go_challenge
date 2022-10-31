package omdb

import (
	"azarc.io/internal/argparser"
	"azarc.io/internal/imdb"
	"strings"
	"sync"
)

type Consumer struct {
	InChannel  chan string
	OutChannel chan *imdb.Movie
	Quit       chan bool
	Workers    int
	Repo       imdb.MovieRepo
	Header     *argparser.TsvHeader
	Config     *argparser.Configuration
}

func (c Consumer) Consume() {
	var wg sync.WaitGroup
	for i := 0; i < c.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case in := <-c.InChannel:
					err := c.worker(in)
					if err != nil {
						return
					}
				case <-c.Quit:
					return
				}
			}
		}()
	}
	wg.Wait()
	c.Quit <- true
	close(c.OutChannel)
}

func (c Consumer) worker(in string) error {
	parts := strings.Split(in, "\t")
	filtered := c.Config.Filters.Filter(Header, parts)
	if filtered {
		movie, err := c.Repo.Retrieve(parts[0])
		if err != nil {
			return err
		}

		if c.Config.RegExpFilter.MatchString(movie.Plot) {
			c.OutChannel <- movie
			return nil
		}

	}
	return nil
}
