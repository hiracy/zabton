package logger

import (
	"log"

	"github.com/comail/colog"
)

func init() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime,
	})
	colog.Register()
}

// Log provides general log output.
func Log(severity, message string) {
	switch severity {
	case "trace", "debug", "info", "warn", "error", "alert":
	case "fatal":
		severity = "error"
	default:
		severity = ""
	}
	log.Printf(severity + ": " + message)
}

func Debug(message string) {
	Log("debug", message)
}
