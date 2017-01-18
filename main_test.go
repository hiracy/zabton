package main

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestVersionFileFormat(t *testing.T) {
	str, err := ioutil.ReadFile("VERSION")
	if err != nil {
		t.Errorf("%v", err)
	}

	matched, err := regexp.MatchString("^[0-9]+.[0-9]+.[0-9]+", string(str))
	if err != nil {
		t.Errorf("%v", err)
	} else if !matched {
		t.Errorf("VERSION file should n.n.n format: %s", string(str))
	}
}
