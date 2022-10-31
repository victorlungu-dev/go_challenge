package main

import (
	"azarc.io/internal/argparser"
	"azarc.io/internal/imdb"
	"azarc.io/internal/imdb/http"
	"azarc.io/internal/omdb"
	"azarc.io/internal/printer"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Shutdown struct {
	Quit            chan bool
	FastQuit        chan bool
	NumOfGoRoutines int
	Termination     chan os.Signal
	Ctx             context.Context
}

func (s Shutdown) Wait() {
	for {
		select {
		case <-s.Ctx.Done():
			fmt.Println("Context exceeded gracefull stop execution")
			s.notifyQuit(s.Quit)
			return
		case <-s.Termination:
			fmt.Println("Received termination signal")
			s.notifyQuit(s.FastQuit)
			return
		}
	}
}

func (s Shutdown) notifyQuit(q chan bool) {
	for i := 0; i < s.NumOfGoRoutines+5; i++ {
		q <- true
	}
}

func initContext(t int) (context.Context, context.CancelFunc) {
	var ctx context.Context
	var cancel context.CancelFunc

	if t > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	} else {
		ctx = context.Background()
	}
	return ctx, cancel
}

func main() {

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	parser := argparser.NewArgParser("IMDB", "IMDB filter and configuration params parser")
	config, err := argparser.ParseIntoConfig(parser, os.Args)
	if err != nil {
		return
	}

	ctx, cancel := initContext(config.Opts.MaxRunTime)
	defer cancel()

	quit := make(chan bool)
	fastQuit := make(chan bool)

	shutdown := Shutdown{
		FastQuit:        fastQuit,
		Quit:            quit,
		NumOfGoRoutines: config.Opts.NumGoRoutines,
		Termination:     termChan,
		Ctx:             ctx,
	}
	go shutdown.Wait()

	// Start line reader that pushes each line on a out channel to be processed by the workers
	out := make(chan string)
	r := omdb.Reader{
		Out:      out,
		Quit:     quit,
		FastQuit: fastQuit,
	}
	go r.StreamLines(config.Opts.FilePath)

	// init movie repository
	repo := http.NewHttpRepo(ctx, "https://www.omdbapi.com/", config.ApiKey, config.Opts.MaxRequests)

	movies := make(chan *imdb.Movie)
	consumer := omdb.Consumer{
		InChannel:  out,
		OutChannel: movies,
		Quit:       quit,
		Workers:    config.Opts.NumGoRoutines,
		Repo:       repo,
		Header:     omdb.Header,
		Config:     config,
	}
	go consumer.Consume()
	p := printer.Printer{W: os.Stdout}
	p.PrintResults(movies, quit, fastQuit)
}
