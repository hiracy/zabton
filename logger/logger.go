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

// SeLevel set log output level.
func SetLevel(level string) {
	switch level {
	case "trace":
		colog.SetMinLevel(colog.LTrace)
	case "debug":
		colog.SetMinLevel(colog.LDebug)
	case "info":
		colog.SetMinLevel(colog.LInfo)
	case "warn":
		colog.SetMinLevel(colog.LWarning)
	case "fatal", "error":
		colog.SetMinLevel(colog.LError)
	case "alert":
		colog.SetMinLevel(colog.LAlert)
	default:
		colog.SetMinLevel(colog.LWarning)
	}
}

// Debug provides debug string.
func Debug(message string) {
	Log("debug", message)
}
