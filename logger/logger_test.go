package logger

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/comail/colog"
)

func TestLog(t *testing.T) {
	buf := new(bytes.Buffer)
	colog.SetOutput(buf)

	SetLevel("fatal")
	Log("warn", "This is warn log.")

	if buf.String() != "" {
		t.Errorf("Fatal level should not be displayed at alert level: %s", buf.String())
	}

	buf.Reset()
	Log("alert", "This is alert log.")

	matched, err := regexp.MatchString(`\[ alert \] [0-9]{4}\/[0-9]{2}\/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} This is alert log\.`, buf.String())
	if err != nil {
		t.Errorf("%v", err)
	} else if !matched {
		t.Errorf("Alert level should display at formatted alert log: %s", buf.String())
	}
}
