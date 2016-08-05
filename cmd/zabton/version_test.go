package main

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestCmdVersion(t *testing.T) {
	str, err := ioutil.ReadFile("../../VERSION")
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = regexp.MatchString("[0-9]+.[0-9]+.[0-9]+", string(str))
	if err != nil {
		t.Errorf("%v", err)
	}
}
