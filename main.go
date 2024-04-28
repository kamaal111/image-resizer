package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	flags, err := initializeFlags()
	if err != nil {
		log.Fatalln(err.Error())
	}

	inputImage, err := openAndReadImage(flags.Input)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(inputImage.Bounds().Max)

	elapsed := time.Since(start)
	fmt.Printf("done resizing image in %s ✨✨✨\n", elapsed)
}

func openAndReadImage(filePath string) (image.Image, error) {
	fileContent, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	fileExtension := extractFileExtension(filePath)
	switch fileExtension {
	case "jpeg", "jpg":
		return jpeg.Decode(fileContent)
	case "png":
		return png.Decode(fileContent)
	default:
		return nil, fmt.Errorf("%s file extension are not supported", fileExtension)
	}
}

func extractFileExtension(filePath string) string {
	filePathSplitByDots := strings.Split(filePath, ".")
	fileExtension := filePathSplitByDots[len(filePathSplitByDots)-1]
	return fileExtension
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
