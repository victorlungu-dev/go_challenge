package main

import (
	"azarc.io/internal/argparser"
	"azarc.io/internal/omdb"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Shutdown struct {
	Quit            chan bool
	NumOfGoRoutines int
	Termination     chan os.Signal
	Ctx             context.Context
}

func (s Shutdown) Stop() {
	for {
		select {
		case <-s.Ctx.Done():
			fmt.Println("Context exceeded stop execution")
			s.notify()
			return
		case <-s.Termination:
			fmt.Println("Received termination signal")
			s.notify()
			return
		}
	}
}

func (s Shutdown) notify() {
	for i := 1; i <= 10; i++ {
		s.Quit <- true
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
	fmt.Println("MAIN")
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	parser := argparser.NewArgParser("IMDB", "IMDB filter and configuration params parser")
	config, err := argparser.ParseIntoConfig(parser, os.Args)
	if err != nil {
		panic(err)
	}
	quit := make(chan bool, 1+config.Opts.NumGoRoutines)

	ctx, cancel := initContext(config.Opts.MaxRunTime)
	defer cancel()

	shutdown := Shutdown{
		Quit:            quit,
		NumOfGoRoutines: 1 + config.Opts.NumGoRoutines,
		Termination:     termChan,
		Ctx:             ctx,
	}
	go shutdown.Stop()

	linesCh := make(chan string, config.Opts.NumGoRoutines)
	producer := omdb.Producer{
		OutChannel: linesCh,
		Quit:       quit,
	}
	go producer.Produce(config.Opts.FilePath)
	movies := make(chan interface{})
	consumer := omdb.Consumer{
		InChannel:     linesCh,
		OutChannel:    movies,
		Quit:          quit,
		NumGoRoutines: config.Opts.NumGoRoutines,
		Filters:       config.Filters,
	}
	consumer.Consume()
	<-quit
	fmt.Println("Closing....")
}
