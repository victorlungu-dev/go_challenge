package omdb

import (
	"bufio"
	"os"
	"strings"
)

type Producer struct {
	OutChannel chan string
	Quit       chan bool
}

func (p Producer) Produce(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	p.scanFile(f)
}

func (p Producer) scanFile(f *os.File) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		select {
		case <-p.Quit:
			close(p.OutChannel)
			return
		case p.OutChannel <- strings.Trim(s.Text(), " "):
		}
	}
	p.Quit <- true
}
