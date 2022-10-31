package argparser_test

import (
	"azarc.io/internal/argparser"
	"fmt"
	"github.com/akamensky/argparse"

	"os"
	"testing"
)

func TestNoArgs(t *testing.T) {
	p := argparser.NewArgParser("test", "test parser")
	_, err := argparser.ParseIntoConfig(p, []string{})
	if err == nil {
		t.Error("Parser should fail with 0 arguments")
	}
}

func TestNewArgParser(t *testing.T) {
	p := argparser.NewArgParser("test", "test parser")
	if p.GetName() != "test" {
		t.Errorf("Invalid parser name")
	}
}

func TestRequiredArguments(t *testing.T) {
	p := argparser.NewArgParser("test", "test parser")
	expArguments := map[string]interface{}{
		"filePath":       "/file/path.tsv",
		"maxRequests":    1,
		"maxRunTime":     2,
		"numGoRoutines":  3,
		"titleType":      "testTitle",
		"primaryTitle":   "testPrimaryTitle",
		"originalTitle":  "testOriginalTitle",
		"genre":          "testGenre",
		"genres":         "testGenres",
		"plotFilter":     "testPlotFilter",
		"startYear":      "1990",
		"endYear":        "1992",
		"runTimeMinutes": "20",
		"apiKey":         "test_api_key",
	}

	config, _ := argparser.ParseIntoConfig(p, []string{"bin",
		"--filePath", fmt.Sprintf("%s", expArguments["filePath"]),
		"--maxRequests", fmt.Sprintf("%d", expArguments["maxRequests"]),
		"--maxRunTime", fmt.Sprintf("%d", expArguments["maxRunTime"]),
		"--numGoRoutines", fmt.Sprintf("%d", expArguments["numGoRoutines"]),
		"--titleType", fmt.Sprintf("%s", expArguments["titleType"]),
		"--primaryTitle", fmt.Sprintf("%s", expArguments["primaryTitle"]),
		"--originalTitle", fmt.Sprintf("%s", expArguments["originalTitle"]),
		"--genre", fmt.Sprintf("%s", expArguments["genre"]),
		"--genres", fmt.Sprintf("%s", expArguments["genres"]),
		"--plotFilter", fmt.Sprintf("%s", expArguments["plotFilter"]),
		"--startYear", fmt.Sprintf("%s", expArguments["startYear"]),
		"--endYear", fmt.Sprintf("%s", expArguments["endYear"]),
		"--runTimeMinutes", fmt.Sprintf("%s", expArguments["runTimeMinutes"]),
		"--apiKey", fmt.Sprintf("%s", expArguments["apiKey"]),
	})

	if config.Opts.FilePath != expArguments["filePath"] {
		t.Errorf("Expected %s got %s", expArguments["filePath"], config.Opts.FilePath)
	}

	if config.Opts.MaxRequests != expArguments["maxRequests"] {
		t.Errorf("Expected %d got %d", expArguments["maxRequests"], config.Opts.MaxRequests)
	}
	if config.Opts.MaxRunTime != expArguments["maxRunTime"] {
		t.Errorf("Expected %d got %d", expArguments["maxRunTime"], config.Opts.MaxRunTime)
	}
	if config.Opts.NumGoRoutines != expArguments["numGoRoutines"] {
		t.Errorf("Expected %d got %d", expArguments["numGoRoutines"], config.Opts.NumGoRoutines)
	}

	if len(config.Filters) != 8 {
		t.Errorf("Expected %d filter got %d", 8, len(config.Filters))
	}

	//if config.RegExpFilter != expArguments["plotFilter"] {
	//	t.Errorf("Expected %s filter got %s", expArguments["plotFilter"], config.RegExpFilter)
	//}

	if config.ApiKey != expArguments["apiKey"] {
		t.Errorf("Expected %s got %s", expArguments["apiKey"], config.ApiKey)
	}

}

func TestApiKeyFromEnv(t *testing.T) {
	p := argparse.NewParser("test", "test parser")
	expArguments := map[string]string{
		"filePath": "/file/path.tsv",
		"apiKey":   "apiKey",
	}
	os.Setenv("OMDB_APIKEY", expArguments["apiKey"])

	config, _ := argparser.ParseIntoConfig(p, []string{"bin",
		"--filePath", expArguments["filePath"],
	})

	if config.ApiKey != expArguments["apiKey"] {
		t.Errorf("Expected %s apiKey got %s", expArguments["apiKey"], config.ApiKey)
	}
}

func TestUnknownArgs(t *testing.T) {
	p := argparse.NewParser("test", "test parser")
	expArguments := map[string]string{
		"filePath":   "/file/path.tsv",
		"invalidArg": "invalid",
	}
	_, err := argparser.ParseIntoConfig(p, []string{"bin",
		"--filePath", expArguments["filePath"],
		"--invalidArg", expArguments["invalidArg"],
	})
	if err == nil {
		t.Errorf("accepted invalid argument")
	}
}
