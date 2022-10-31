package printer

import (
	"azarc.io/internal/imdb"
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	printer := Printer{W: os.Stdout}
	movies := []*imdb.Movie{
		{Title: "TestTitle", Plot: "testPlot", ImdbID: "testId"},
		{Title: "TestTitle1", Plot: "testPlot", ImdbID: "testId1"},
		{Title: "TestTitle2", Plot: "testPlot", ImdbID: "testId2"},
		{Title: "TestTitle3", Plot: "testPlot", ImdbID: "testId3"},
		{Title: "TestTitle4", Plot: "testPlot", ImdbID: "testId4"},
	}
	printer.PrettyPrint(movies)
}

func TestPrintResultsWithQuit(t *testing.T) {
	var output bytes.Buffer
	printer := Printer{W: &output}
	in := make(chan *imdb.Movie, 1)
	quit := make(chan bool, 1)
	fastQuit := make(chan bool, 1)
	go printer.PrintResults(in, quit, fastQuit)
	in <- &imdb.Movie{Title: "testTitle", Plot: "Test plot", ImdbID: "testId"}
	if !strings.Contains("testTile", output.String()) {
		t.Errorf("expected line in output %s", "testTitle")
	}
	quit <- true
}
