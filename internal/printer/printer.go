package printer

import (
	"azarc.io/internal/imdb"
	"fmt"
	"io"
	"text/tabwriter"
)

type Printer struct {
	W io.Writer
}

func (p Printer) PrettyPrint(movies []*imdb.Movie) {
	if len(movies) > 0 {
		w := tabwriter.NewWriter(p.W, 0, 0, 3, ' ', tabwriter.Debug)
		fmt.Fprintln(w, fmt.Sprintf("IMDB_ID\t Title\t Plot\t"))
		for _, m := range movies {
			fmt.Fprintln(w, m.String())
		}
		w.Flush()
	}
}

func (p Printer) PrintResults(in chan *imdb.Movie, quit chan bool, fastQuit chan bool) {
	var movies []*imdb.Movie
	for {
		select {
		case <-fastQuit:
			return
		case movie := <-in:
			if movie != nil && movie.ImdbID != "" {
				movies = append(movies, movie)
			}
		case <-quit:
			p.PrettyPrint(movies)
			return
		}
	}
}
