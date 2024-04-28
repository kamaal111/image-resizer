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
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/draw"
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

	_, err = createResizedImage(inputImage, flags.Output, flags.Dimensions)
	if err != nil {
		log.Fatalln(err.Error())
	}

	elapsed := time.Since(start)
	fmt.Printf("done resizing image in %s ✨✨✨\n", elapsed)
}

func createResizedImage(inputImage image.Image, outputPath string, desiredDimensions Dimensions) (*os.File, error) {
	outputImage, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}
	defer outputImage.Close()

	outputImageSpec := image.NewRGBA(image.Rect(0, 0, desiredDimensions.Width, desiredDimensions.Height))
	draw.NearestNeighbor.Scale(outputImageSpec, outputImageSpec.Rect, inputImage, inputImage.Bounds(), draw.Over, nil)
	png.Encode(outputImage, outputImageSpec)

	return outputImage, nil
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

type Dimensions struct {
	Width  int
	Height int
}

type Flags struct {
	Input      string
	Output     string
	Dimensions Dimensions
}

func initializeFlags() (*Flags, error) {
	inputPath := flag.String("i", "", "input path")
	outputPath := flag.String("o", "", "output path")
	dimensions := flag.String("d", "", "dimensions")
	flag.Parse()

	if *outputPath == "" {
		return nil, errors.New("no output path provided\nplease give a output path by giving this command the -o flag with the destination")
	}
	if *inputPath == "" {
		return nil, errors.New("no input path provided\nplease give a input path by giving this command the -i flag with the destination")
	}
	if *dimensions == "" {
		return nil, errors.New("no dimensions provided\nplease provided the wished for dimensions, by passing this command -d flag with the desired dimensions")
	}

	dimensionsSplitByX := strings.Split(strings.ToLower(strings.ReplaceAll(*dimensions, " ", "")), "x")
	if len(dimensionsSplitByX) != 2 {
		return nil, errors.New("invalid dimensions provided\ndimensions should be formatted as '123x123'")
	}

	var dimensionsAsIntegers []int
	for _, dimension := range dimensionsSplitByX {
		dimensionAsInteger, err := strconv.Atoi(dimension)
		if err != nil {
			return nil, err
		}

		dimensionsAsIntegers = append(dimensionsAsIntegers, dimensionAsInteger)
	}

	return &Flags{
		Input:  *inputPath,
		Output: *outputPath,
		Dimensions: Dimensions{
			Width:  dimensionsAsIntegers[0],
			Height: dimensionsAsIntegers[1],
		},
	}, nil
}
