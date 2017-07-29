package golog

import (
	"encoding/json"
	"os"
	"os/user"
	"path"

	"github.com/kataras/pio"
)

// Level is a number which defines the log level.
type Level uint32

// The available log levels.
const (
	// DisableLevel will disable printer
	DisableLevel Level = iota
	// ErrorLevel will print only errors
	ErrorLevel
	// WarnLevel will print errors and warnings
	WarnLevel
	// InfoLevel will print errors, warnings and infos
	InfoLevel
	// DebugLevel will print on any level, errors, warnings, infos and debug messages
	DebugLevel
)

var (
	// without colors
	erroText = choose("[ERRO]", "[ERROR]")
	warnText = choose("[WARN]", "[WARNING]")
	infoText = choose("[INFO]", "[INFORMATION]")
	dbugText = choose("[DBUG]", "[DEBUG]")
	// with colors
	erro = pio.Red(erroText)
	warn = pio.Purple(warnText)
	info = pio.LightGreen(infoText)
	dbug = pio.Yellow(dbugText)
)

// Config struct is read from ~/.golog config file
type Config struct {
	UseShortMessages bool
}

func choose(shortMsg, longMsg string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	configFile := path.Join(usr.HomeDir, ".golog")
	file, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	if config.UseShortMessages {
		return shortMsg
	}
	return longMsg
}

// returns a [PREFIX] based on the "level" and "enableColor".
func prefixFromLevel(level Level, enableColor bool) string {
	switch level {
	case ErrorLevel:
		if !enableColor {
			return erroText
		}
		return erro
	case WarnLevel:
		if !enableColor {
			return warnText
		}
		return warn
	case InfoLevel:
		if !enableColor {
			return infoText
		}
		return info
	case DebugLevel:
		if !enableColor {
			return dbugText
		}
		return dbug
	default:
		return ""
	}
}

func fromLevelName(levelName string) Level {
	switch levelName {
	case "error":
		return ErrorLevel
	case "warning":
		fallthrough
	case "warn":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	default:
		return DisableLevel
	}
}
