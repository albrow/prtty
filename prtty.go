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
	defaultColor = "@w" // white
	infoColor    = "@c" // cyan
	warnColor    = "@y" // yellow
	successColor = "@g" // green
	errorColor   = "@r" // red
)

var (
	Default = NewLogger(os.Stdout, defaultColor)
	Info    = NewLogger(os.Stdout, infoColor)
	Warn    = NewLogger(os.Stdout, warnColor)
	Success = NewLogger(os.Stdout, successColor)
	Error   = NewLogger(os.Stderr, errorColor)
)

type Logger struct {
	Output io.Writer
	Color  string
}

func NewLogger(out io.Writer, color string) *Logger {
	return &Logger{
		Output: out,
		Color:  color,
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
