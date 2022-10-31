package imdb

type MovieRepo interface {
	Retrieve(id string) (*Movie, error)
}
