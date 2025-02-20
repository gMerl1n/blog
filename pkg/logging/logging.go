package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/gMerl1n/blog/internal/config"
	"github.com/sirupsen/logrus"
	l "github.com/sirupsen/logrus"
)

func InitLogger(cfg *config.ConfigServer) *l.Logger {

	log := l.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0755)
	if err != nil || os.IsExist(err) {
		panic("can't create log dir. no configured logging to files")
	} else {
		allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(fmt.Sprintf("[Error]: %s", err))
		}

		mw := io.MultiWriter(os.Stdout, allFile) // MultiWriter to log to stdout and file

		log.SetOutput(mw) // Send all logs to nowhere by default

	}

	log.Formatter = &l.JSONFormatter{}
	log.Level = l.Level(cfg.LogLevel)

	return log
}
