package argparser

import (
	"errors"
	"fmt"
	arglib "github.com/akamensky/argparse"
)

type Options struct {
	FilePath      string // absolute path to the file to be processed
	MaxApiRequest string // maximum number of requests to be made to the omdbapi
	MaxRunTime    int    //
	MaxRequests   string //
	NumGoRoutines int    //
}

// Needs a refactor with a struct filter and handling the list
type Filters struct {
	TitleType      string //
	PrimaryTitle   string //
	OriginalTitle  string //
	Genre          string //
	IsAdult        string //
	StartYear      int    //
	EndYear        int    //
	RunTimeMinutes int    //
	Genres         string //
	PlotFilter     string //
}

type Configuration struct {
	Opts    Options
	Filters Filters
}

func NewArgParser(name, decription string) *arglib.Parser {
	parser := arglib.NewParser(name, decription)
	return parser
}

func ParseIntoConfig(p *arglib.Parser, args []string) (*Configuration, error) {
	if len(args) == 0 {
		msg := "not enough arguments"
		fmt.Print(p.Usage(msg))
		return nil, errors.New(msg)
	}

	filePath := p.String("", "filePath", &arglib.Options{Required: true, Help: "String to print"})
	maxRequests := p.String("", "maxRequests", &arglib.Options{Help: "String to print"})
	maxRunTime := p.Int("", "maxRunTime", &arglib.Options{Help: "String to print"})
	maxApiRequest := p.String("", "maxApiRequest", &arglib.Options{Help: "String to print"})
	numGoRoutines := p.Int("", "numGoRoutines", &arglib.Options{Help: "String to print"})

	titleType := p.String("", "titleType", &arglib.Options{Help: "String to print"})
	primaryTitle := p.String("", "primaryTitle", &arglib.Options{Help: "String to print"})
	originalTitle := p.String("", "originalTitle", &arglib.Options{Help: "String to print"})
	isAdult := p.String("", "isAdult", &arglib.Options{Help: "String to print"})
	genre := p.String("", "genre", &arglib.Options{Help: "String to print"})
	plotFilter := p.String("", "plotFilter", &arglib.Options{Help: "String to print"})
	startYear := p.Int("", "startYear", &arglib.Options{Help: "String to print"})
	endYear := p.Int("", "endYear", &arglib.Options{Help: "String to print"})
	genres := p.String("", "genres", &arglib.Options{Help: "String to print"})
	runTimeMinutes := p.Int("", "runTimeMinutes", &arglib.Options{Help: "String to print"})

	err := p.Parse(args)

	if err != nil {
		// In case of error print error and print usage and return it
		// This can also be done by passing -h or --help flags
		fmt.Print(p.Usage(err))
		return nil, err
	}
	config := Configuration{
		Opts: Options{
			FilePath:      *filePath,
			MaxApiRequest: *maxApiRequest,
			MaxRunTime:    *maxRunTime,
			MaxRequests:   *maxRequests,
			NumGoRoutines: *numGoRoutines,
		},
		Filters: Filters{
			TitleType:      *titleType,
			PrimaryTitle:   *primaryTitle,
			OriginalTitle:  *originalTitle,
			IsAdult:        *isAdult,
			Genre:          *genre,
			StartYear:      *startYear,
			EndYear:        *endYear,
			RunTimeMinutes: *runTimeMinutes,
			Genres:         *genres,
			PlotFilter:     *plotFilter,
		},
	}
	return &config, nil
}
