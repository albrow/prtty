// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

// package prtty is a small go library for logging things with color.
package prtty

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"io"
	"log"
	"os"
)

const (
	// Color settings for the different logging functions
	// See: https://godoc.org/github.com/wsxiaoys/terminal/color
	ColorRed     = "@r"
	ColorGreen   = "@g"
	ColorBlue    = "@b"
	ColorCyan    = "@c"
	ColorMagenta = "@m"
	ColorYellow  = "@y"
	ColorBlack   = "@k"
	ColorWhite   = "@w"
)

var (
	Default        = NewLogger(os.Stdout, ColorWhite)
	Info           = NewLogger(os.Stdout, ColorCyan)
	Warn           = NewLogger(os.Stdout, ColorYellow)
	Success        = NewLogger(os.Stdout, ColorGreen)
	Error          = NewLogger(os.Stderr, ColorRed)
	DefaultLoggers LoggerGroup
	AllLoggers     LoggerGroup
)

type Logger struct {
	Output io.Writer
	Color  string
}

type LoggerGroup []*Logger

func init() {
	DefaultLoggers = LoggerGroup{Default, Info, Warn, Success, Error}
	AllLoggers = append(LoggerGroup{}, DefaultLoggers...)
}

func NewLogger(out io.Writer, color string) *Logger {
	logger := &Logger{
		Output: out,
		Color:  color,
	}
	AllLoggers = append(AllLoggers, logger)
	return logger
}

func (lg LoggerGroup) SetOutput(out io.Writer) {
	for _, logger := range lg {
		logger.Output = out
	}
}

func (lg LoggerGroup) SetColor(color string) {
	for _, logger := range lg {
		logger.Color = color
	}
}

func (l *Logger) Print(v ...interface{}) {
	log.SetOutput(l.Output)
	log.Print(color.Sprint(l.Color + fmt.Sprint(v...)))
}

func (l *Logger) Println(v ...interface{}) {
	log.SetOutput(l.Output)
	log.Println(color.Sprint(l.Color + fmt.Sprint(v...)))
}

func (l *Logger) Printf(format string, v ...interface{}) {
	log.SetOutput(l.Output)
	log.Printf(color.Sprint(l.Color + fmt.Sprintf(format, v...)))
}

func (l *Logger) Panic(v ...interface{}) {
	log.SetOutput(l.Output)
	log.Panic(color.Sprint(l.Color + fmt.Sprint(v...)))
}

func (l *Logger) Panicln(v ...interface{}) {
	log.SetOutput(l.Output)
	log.Panicln(color.Sprint(l.Color + fmt.Sprint(v...)))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	log.SetOutput(l.Output)
	log.Panicf(color.Sprint(l.Color + fmt.Sprintf(format, v...)))
}

func (l *Logger) Fatal(v ...interface{}) {
	log.SetOutput(l.Output)
	log.Fatal(color.Sprint(l.Color + fmt.Sprint(v...)))
}

func (l *Logger) Fatalln(v ...interface{}) {
	log.SetOutput(l.Output)
	log.Fatalln(color.Sprint(l.Color + fmt.Sprint(v...)))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	log.SetOutput(l.Output)
	log.Fatalf(color.Sprint(l.Color + fmt.Sprintf(format, v...)))
}
