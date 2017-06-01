package main

import (
	"os/exec"
	"regexp"
	"testing"
)

func TestVersionFormat(t *testing.T) {
	str, err := exec.Command("git", "tag").Output()
	if err != nil {
		t.Errorf("%v", err)
	}

	matched, err := regexp.MatchString(`^v[0-9]+.[0-9]+.[0-9]+`, string(str))
	if err != nil {
		t.Errorf("%v", err)
	} else if !matched {
		t.Errorf("vertion(git tag output) should vn.n.n format: %s", string(str))
	}
}
