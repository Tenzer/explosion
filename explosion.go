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

var (
	height uint
	width  uint
)

func PrintImage(img image.Image) {
	resized := resize.Thumbnail(width, height, img, resize.NearestNeighbor)
	bounds := resized.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			upper_red, upper_green, upper_blue, _ := resized.At(x, y).RGBA()
			lower_red, lower_green, lower_blue, _ := resized.At(x, y+1).RGBA()

			// Only print the background if upper and lower row are same color
			if upper_red == lower_red && upper_green == lower_green && upper_blue == lower_blue {
				fmt.Printf(
					"\x1b[48;2;%d;%d;%dm ",
					upper_red/256,
					upper_green/256,
					upper_blue/256,
				)
			} else {
				fmt.Printf(
					"\x1b[48;2;%d;%d;%dm\x1b[38;2;%d;%d;%dm▄",
					upper_red/256,
					upper_green/256,
					upper_blue/256,
					lower_red/256,
					lower_green/256,
					lower_blue/256,
				)
			}
		}
		fmt.Println("\x1b[0m")
	}
}

func main() {
	height_int, width_int, _ := pty.Getsize(os.Stdout)

	// The three subtracted lines is to have room for command, file name and prompt after explosion
	flag.UintVar(&height, "h", uint((height_int-3)*2), "Maximum height of output in number of half lines")
	flag.UintVar(&width, "w", uint(width_int), "Maximum width of output in number of columns")
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file ...]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "  Specify \"-\" to read from stdin.\n")
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

		source_image, _, err := image.Decode(file)
		file.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			continue
		}

		PrintImage(source_image)
	}
}
