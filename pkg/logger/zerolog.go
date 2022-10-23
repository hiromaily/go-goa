package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/hiromaily/go-goa/pkg/config"
)

// NewZeroLog configures globally
// - https://github.com/rs/zerolog
func InitializeZeroLog(conf *config.Logger) {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLv := zerolog.Level(conf.Level)
	if logLv.String() == "" {
		logLv = zerolog.Disabled
	}
	zerolog.SetGlobalLevel(logLv)

	if conf.FileName != "" {
		logFile, err := os.OpenFile(conf.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
		if err != nil {
			panic(err)
		}
		log.Logger = log.Output(logFile)

		// FIXME: where to close
		logFile.Close()
	}
}
