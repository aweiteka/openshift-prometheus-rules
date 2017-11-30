package rules_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/prometheus/prometheus/promql"
)

var testdir string

func init() {
	flag.StringVar(&testdir, "testdir", "data", "Path to testdir")
}

func TestRules(t *testing.T) {
	flag.Parse()
	searchpath := fmt.Sprintf("%s/*.txt", testdir)
	fmt.Println("Collecting data files from", searchpath)
	files, err := filepath.Glob(searchpath)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		fmt.Println("Evaluating data file", file)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			t.Error(err)
			continue
		}
		test, err := promql.NewTest(t, string(data))
		if err != nil {
			t.Errorf("Failed to create test for %s: %s", file, err)
		}
		test.Run()
		if err != nil {
			t.Errorf("Failed to run test %s: %s", file, err)
		}
		test.Close()
	}
}
