package argparser

import (
	"errors"
	"fmt"
	arglib "github.com/akamensky/argparse"
	"os"
	"regexp"
)

type Options struct {
	FilePath      string
	MaxRunTime    int
	MaxRequests   int
	NumGoRoutines int
}

type Configuration struct {
	Opts         Options
	Filters      StringFilters
	RegExpFilter *regexp.Regexp
	ApiKey       string
}

func NewArgParser(name, description string) *arglib.Parser {
	parser := arglib.NewParser(name, description)
	return parser
}

func ParseIntoConfig(p *arglib.Parser, args []string) (*Configuration, error) {
	if len(args) == 0 {
		msg := "not enough arguments"
		fmt.Print(p.Usage(msg))
		return nil, errors.New(msg)
	}
	// configurations
	filePath := p.String("", "filePath", &arglib.Options{Required: true, Help: "String to print"})
	apiKey := p.String("", "apiKey", &arglib.Options{Help: "Omdb api key"})
	maxRequests := p.Int("", "maxRequests", &arglib.Options{Help: "Max number of requests to be made to omdbAPi"})
	maxRunTime := p.Int("", "maxRunTime", &arglib.Options{Help: "Max number of seconds that the program can run before softquitting"})
	//maxApiRequest := p.String("", "maxApiRequest", &arglib.Options{Help: "String to print"})
	numGoRoutines := p.Int("", "numGoRoutines", &arglib.Options{Help: "Number of goroutines to use for the task", Default: 10})

	// filters
	titleType := p.String("", "titleType", &arglib.Options{Help: "Title filter"})
	primaryTitle := p.String("", "primaryTitle", &arglib.Options{Help: "Primary title filter"})
	originalTitle := p.String("", "originalTitle", &arglib.Options{Help: "Original title filter"})
	genre := p.String("", "genre", &arglib.Options{Help: "Genre filter"})
	genres := p.String("", "genres", &arglib.Options{Help: "Genres filter"})
	plotFilter := p.String("", "plotFilter", &arglib.Options{Help: "Regexp to us in plot filtering from omdbAPi"})

	startYear := p.String("", "startYear", &arglib.Options{Help: "Start year filter"})
	runTimeMinutes := p.String("", "runTimeMinutes", &arglib.Options{Help: "Runtimeminutes filter"})
	endYear := p.String("", "endYear", &arglib.Options{Help: "End year filter"})

	err := p.Parse(args)

	if err != nil {
		// In case of error print error and print usage and return it
		// This can also be done by passing -h or --help flags
		fmt.Print(p.Usage(err))
		return nil, err
	}

	stringFilters := StringFilters{}
	if *startYear != "" {
		stringFilters = append(stringFilters, NewStringFilter("startYear", *startYear))
	}
	if *endYear != "" {
		stringFilters = append(stringFilters, NewStringFilter("endYear", *endYear))
	}
	if *runTimeMinutes != "" {
		stringFilters = append(stringFilters, NewStringFilter("runTimeMinutes", *runTimeMinutes))
	}
	if *titleType != "" {
		stringFilters = append(stringFilters, NewStringFilter("titleType", *titleType))
	}

	if *primaryTitle != "" {
		stringFilters = append(stringFilters, NewStringFilter("primaryTitle", *primaryTitle))
	}

	if *originalTitle != "" {
		stringFilters = append(stringFilters, NewStringFilter("originalTitle", *originalTitle))
	}

	if *genre != "" {
		stringFilters = append(stringFilters, NewStringFilter("genre", *genre))
	}

	if *genres != "" {
		stringFilters = append(stringFilters, NewStringFilter("genres", *genres))
	}
	if *apiKey == "" {
		*apiKey = os.Getenv("OMDB_APIKEY")
	}

	fmt.Println(*plotFilter)
	re := regexp.MustCompile(*plotFilter)
	config := Configuration{
		Opts: Options{
			FilePath:      *filePath,
			MaxRunTime:    *maxRunTime,
			MaxRequests:   *maxRequests,
			NumGoRoutines: *numGoRoutines,
		},
		Filters:      stringFilters,
		RegExpFilter: re,
		ApiKey:       *apiKey,
	}
	return &config, nil
}
