package logger

import (
	"fmt"
	"os"

	"github.com/hiromaily/go-goa/pkg/config"
	"github.com/rs/zerolog/log"
)

// logger interface
//  stdout
//  local file
//  pubsub on GCP

type Logger interface {
	Printf(format string, v ...any)
	Print(v ...any)
	Println(v ...any)
	Fatalf(format string, v ...any)
	Fatal(v ...any)
	Fatalln(v ...any)
}

type zeroLog struct{}

func NewZeroLog(conf *config.Logger) Logger {
	InitializeZeroLog(conf)
	return &zeroLog{}
}

func (z *zeroLog) Printf(format string, v ...any) {
	log.Info().Msg(fmt.Sprintf(format, v...))
}

func (z *zeroLog) Print(v ...any) {
	log.Info().Msg(fmt.Sprint(v...))
}

func (z *zeroLog) Println(v ...any) {
	log.Info().Msg(fmt.Sprint(v...))
}

func (z *zeroLog) Fatalf(format string, v ...any) {
	log.Info().Msg(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (z *zeroLog) Fatal(v ...any) {
	log.Info().Msg(fmt.Sprint(v...))
	os.Exit(1)
}

func (z *zeroLog) Fatalln(v ...any) {
	log.Info().Msg(fmt.Sprint(v...))
	os.Exit(1)
}
