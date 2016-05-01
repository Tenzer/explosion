package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/kr/pty"
	"github.com/nfnt/resize"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func PrintImage(img image.Image, width uint, height uint) {
	resized := resize.Thumbnail(width, height, img, resize.NearestNeighbor)
	bounds := resized.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			upperRed, upperGreen, upperBlue, _ := resized.At(x, y).RGBA()
			lowerRed, lowerGreen, lowerBlue, _ := resized.At(x, y+1).RGBA()

			// Only print the background if upper and lower row are same color
			if upperRed == lowerRed && upperGreen == lowerGreen && upperBlue == lowerBlue {
				fmt.Printf(
					"\x1b[48;2;%d;%d;%dm ",
					upperRed/256,
					upperGreen/256,
					upperBlue/256,
				)
			} else {
				fmt.Printf(
					"\x1b[48;2;%d;%d;%dm\x1b[38;2;%d;%d;%dm▄",
					upperRed/256,
					upperGreen/256,
					upperBlue/256,
					lowerRed/256,
					lowerGreen/256,
					lowerBlue/256,
				)
			}
		}
		fmt.Println("\x1b[0m")
	}
}

func main() {
	heightInt, widthInt, _ := pty.Getsize(os.Stdout)

	var width uint
	var height uint

	// The three subtracted lines is to have room for command, file name and prompt after explosion
	flag.UintVar(&width, "w", uint(widthInt), "Maximum width of output in number of columns")
	flag.UintVar(&height, "h", uint((heightInt-3)*2), "Maximum height of output in number of half lines")
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file ...]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "  Specify \"-\" to read from stdin.")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	filenames := flag.Args()
	if len(filenames) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for i, filename := range filenames {
		if i > 0 {
			fmt.Println()
		}

		var file *os.File
		var err error

		if filename == "-" {
			fmt.Println("stdin:")
			file = os.Stdin
		} else {
			fmt.Printf("%s:\n", filename)
			file, err = os.Open(filename)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				continue
			}
		}

		sourceImage, _, err := image.Decode(file)
		file.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			continue
		}

		PrintImage(sourceImage, width, height)
	}
}
