package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	flags, err := initializeFlags()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(flags.Input, flags.Output)

	elapsed := time.Since(start)
	fmt.Printf("done resizing image in %s\n", elapsed)
}

type Flags struct {
	Input  string
	Output string
}

func initializeFlags() (*Flags, error) {
	inputPath := flag.String("i", "", "input path")
	outputPath := flag.String("o", "", "output path")
	flag.Parse()
	if *outputPath == "" {
		return nil, errors.New("no output path provided\nplease give a output path by giving this command the -o flag with the destination")
	}
	if *inputPath == "" {
		return nil, errors.New("no input path provided\nplease give a input path by giving this command the -i flag with the destination")
	}

	return &Flags{
		Input:  *inputPath,
		Output: *outputPath,
	}, nil
}
