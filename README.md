# prtty

[![GoDoc](https://godoc.org/github.com/albrow/prtty?status.svg)](https://godoc.org/github.com/albrow/prtty)

A small go library for logging things with color.

### Basics

There are five different loggers includeded out of the box:

- Default (white)
- Info    (cyan)
- Warn    (yellow)
- Success (green)
- Error   (red)

All the included loggers print to os.Stdout, except for Error, which prints to os.Stderr.

You can also create custom loggers with [prtty.NewLogger](http://godoc.org/github.com/albrow/prtty#NewLogger).

The [Logger](http://godoc.org/github.com/albrow/prtty#Logger) type has most of the same methods as the builtin
`log` package. In fact, it is just a lightweight wrapper around it.

prtty respects the flags set with [log.SetFlags](http://golang.org/pkg/log/#SetFlags).

### Usage

``` go
package main

import (
	"github.com/albrow/prtty"
)

func main() {
	// This will print to the terminal in cyan
	prtty.Info.Println("--> starting")

	// This will print in white
	prtty.Default.Println("    doing one thing...")
	prtty.Default.Println("    doing another thing...")

	// This will print in red, and just like log.Fatal, will cause the program
	// to exit with a return code of 1.
	prtty.Error.Fatal("ERROR: something went wrong :(")
}
```

Depending on your terminal color settings, the above program would produce output like this:

![example output](http://oi58.tinypic.com/anfu5d.jpg)

### License

prtty is licensed under the MIT License. See the LICENSE file for more information.
