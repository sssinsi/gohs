package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	count      = flag.Uint("c", 5, "[Option] : Set split line count")
	outputPath = flag.String("o", "", "[Option] : Set output file path")
	path       = flag.String("p", "", "[required] : Set source file path")
	splitter   = flag.String("s", "", "[Option] : Set split string(ex: ************, ///////////// )")
)

type outputType int

const (
	std outputType = iota
	file
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addNewLine(o *os.File, line []byte) {
	_, err := o.Write(line)
	check(err)
	addLineFeed(o)
}

func addLineFeed(o *os.File) {
	_, err := o.Write(([]byte)("\n"))
	check(err)
}

func main() {
	var ot outputType = std
	flag.Parse()

	if *count < 1 {
		fmt.Println(" (s>0) : split line count need to be greater than zero")
		os.Exit(1)
	}

	f, err := os.Open(*path)
	check(err)
	defer f.Close()

	var o *os.File
	if *outputPath != "" {
		o, err = os.OpenFile(*outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		check(err)
		ot = file
	}
	defer o.Close()

	reader := bufio.NewReaderSize(f, 4096)
	var lc uint

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		check(err)

		lc++

		switch ot {
		case std:
			fmt.Println(string(line))

			if lc%*count == 0 {
				fmt.Println(*splitter)
			}
		case file:
			addNewLine(o, line)

			if lc%*count == 0 {
				addNewLine(o, ([]byte)(*splitter))
			}
		}

	}

	switch ot {
	case std:
	case file:
		o.Sync()
	}

}
