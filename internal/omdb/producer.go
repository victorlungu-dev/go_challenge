package omdb

import (
	"azarc.io/internal/argparser"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// global header of the input file
var Header *argparser.TsvHeader

type Reader struct {
	Out      chan string
	Quit     chan bool
	FastQuit chan bool
}

func (r Reader) StreamLines(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("invalid path or file doesn't exist")
		r.Quit <- true
		return
	}
	defer f.Close()
	r.scanFile(f)
}

func (r Reader) scanFile(f *os.File) {
	s := bufio.NewScanner(f)
	defer close(r.Out)
	for s.Scan() {
		select {
		case <-r.FastQuit:
			return
		case <-r.Quit:
			return
		default:
			if Header == nil {
				initHeader(strings.Trim(s.Text(), " "))
			} else {
				r.Out <- strings.Trim(s.Text(), " ")
			}
		}
	}
	fmt.Println("Scan complete")
}

func initHeader(ln string) {
	Header = &argparser.TsvHeader{Columns: map[string]int{}}
	parts := strings.Split(ln, "\t")
	for i, v := range parts {
		Header.Columns[v] = i
	}
}
