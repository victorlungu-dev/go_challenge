package omdb

import "testing"

func TestConsumerWorker(t *testing.T) {
	inChan := make(chan string)
	outChan := make(chan interface{})
	quit := make(chan bool)
	numGoroutines := 1
	consumer := Consumer{
		InChannel:     inChan,
		OutChannel:    outChan,
		Quit:          quit,
		NumGoRoutines: numGoroutines,
	}
	go consumer.Consume()
	inChan <- "test string"
	quit <- true
}
