package omdb

import (
	"azarc.io/internal/argparser"
	"fmt"
	"strings"
)

type Consumer struct {
	InChannel     chan string
	OutChannel    chan interface{}
	Quit          chan bool
	NumGoRoutines int
	Filters       argparser.Filters
}

func (c Consumer) Consume() {
	for i := 0; i <= c.NumGoRoutines; i++ {
		go func(workerId int) {
			fmt.Printf("Start worker-%d\n", workerId)
			for {
				select {
				case in := <-c.InChannel:
					c.worker(in)
				case <-c.Quit:
					return
				}
			}
		}(i)
	}

}

func (c Consumer) worker(in string) {

	filtered := c.filter(in)
	if filtered {
		fmt.Println(in)
		//movie := NewMovieFromString(in)
		//c.OutChannel <- movie
	}
}

func (c Consumer) filter(in string) bool {
	// this needs a refactor to apply all the filters
	parts := strings.Split(in, "\t")

	if c.Filters.PrimaryTitle != "" && c.Filters.PrimaryTitle == parts[2] {
		return true
	}
	if c.Filters.PrimaryTitle == "" {
		return true
	}
	return false
}
