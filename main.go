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
	splitter   = flag.String("s", "", "[Option] : Set split string")
	outputPath = flag.String("o", "", "[Option] : Set output file path")
	path       = flag.String("p", "", "[required] : Set source file path")
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

func exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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
		if exist(*outputPath) {
			fmt.Println("exist")
			o, err = os.Open(*outputPath)
			check(err)
		} else {
			o, err = os.Create(*outputPath)
			check(err)
		}
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
			fmt.Println("output file")
			o.Write(line)
			o.Write(([]byte)("\n"))
			if lc%*count == 0 {
				_, err := o.Write(([]byte)(*splitter))
				check(err)
				_, err = o.Write(([]byte)("\n"))
				check(err)
			}
		}

	}
	switch ot {
	case std:
	case file:
		o.Sync()
	}

}
