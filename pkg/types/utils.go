package types

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logFormatter log.Formatter
var LogOutput *lumberjack.Logger //io.Writer
var logLevel log.Level

func ConfigureLogger(clog *log.Logger) error {
	/* Configure logs */

	if LogOutput != nil {
		clog.SetOutput(LogOutput)
	}
	if logFormatter != nil {
		clog.SetFormatter(logFormatter)
	}
	clog.SetLevel(logLevel)
	return nil
}
