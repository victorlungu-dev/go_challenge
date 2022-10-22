package argparser_test

import (
	"azarc.io/internal/argparser"
	"fmt"
	"github.com/akamensky/argparse"
	"testing"
)

func TestNoArgs(t *testing.T) {
	p := argparse.NewParser("test", "test parser")
	_, err := argparser.ParseIntoConfig(p, []string{})
	if err == nil {
		t.Error("Parser should fail with 0 arguments")
	}
}

func TestRequiredArguments(t *testing.T) {
	p := argparse.NewParser("test", "test parser")
	expPath := "/some/file/path"
	expMaxRunTime := 2
	// the first argument from command line is the name of the bin
	config, _ := argparser.ParseIntoConfig(p, []string{"bin", "--filePath", expPath, "--maxRunTime", fmt.Sprintf("%d", expMaxRunTime)})
	if config.Opts.FilePath != expPath {
		t.Errorf("Expected %s got %s", expPath, config.Opts.FilePath)
	}
	fmt.Println(config)

	if config.Opts.MaxRunTime != expMaxRunTime {
		t.Errorf("Expected %d got %s", expMaxRunTime, config.Opts.FilePath)
	}

}
