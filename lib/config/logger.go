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
			l = strings.ToUpper(l)
			switch l {
			case "TRACE":
				l = Colorize(l, Magenta)
			case "DEBUG":
				l = Colorize(l, Green)
			case "INFO":
				l = Colorize(l, Green)
			case "WARN":
				l = Colorize(l, Yellow)
			case "ERROR":
				l = Colorize(Colorize(l, Red), Bold)
			case "FATAL":
				l = Colorize(Colorize(l, Red), Bold)
			case "PANIC":
				l = Colorize(Colorize(l, Red), Bold)
			default:
				l = Colorize(l, Bold)
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
		return fmt.Sprintf("%s = ", i)
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
