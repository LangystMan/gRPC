package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

const (
	Red = iota + 31
	Green
	Yellow

	Magenta = 35
	Bold    = 1
)

func LoadLogger() zerolog.Logger {

	logWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.StampMilli,
	}

	logWriter.FormatLevel = func(i interface{}) string {
		if l, ok := i.(string); ok {
			switch l {
			case "trace":
				l = Colorize("TRC", Magenta)
			case "debug":
				l = Colorize("DBG", Green)
			case "info":
				l = Colorize("INF", Green)
			case "warn":
				l = Colorize("WRN", Yellow)
			case "error":
				l = Colorize(Colorize("ERR", Red), Bold)
			case "fatal":
				l = Colorize(Colorize("FTL", Red), Bold)
			case "panic":
				l = Colorize(Colorize("PNC", Red), Bold)
			default:
				l = Colorize("???", Bold)
			}

			return fmt.Sprintf("| %s |", l)
		} else {
			if i == nil {
				return Colorize("???", Bold)
			} else {
				return fmt.Sprintf("| %s |", Colorize(strings.ToUpper(i.(string)), Bold)[0:3])
			}
		}
	}

	logWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}

	logWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	logger := log.Output(logWriter)

	return logger
}

func Colorize(s interface{}, c int) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}
