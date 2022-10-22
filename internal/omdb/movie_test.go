package omdb

import "testing"

func TestMovieCreation(t *testing.T) {
	part := "tt0000834\tshort\tA Coward's Courage\tA Coward's Courage\t0\t1909\t\\N\t\\N\t\\N"
	movie := NewMovieFromString(part)
	expTconst := "tt0000834"
	if expTconst != movie.Tconst {
		t.Errorf("Expected %s got %s for movie creation", expTconst, movie.Tconst)
	}
}
