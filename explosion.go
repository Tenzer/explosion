package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/nfnt/resize"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func PrintImage(img image.Image) {
	resized := resize.Thumbnail(80, 80, img, resize.NearestNeighbor)
	bounds := resized.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			upper_red, upper_green, upper_blue, _ := resized.At(x, y).RGBA()
			lower_red, lower_green, lower_blue, _ := resized.At(x, y+1).RGBA()
			fmt.Printf("\x1b[48;2;%d;%d;%dm\x1b[38;2;%d;%d;%dm▄", upper_red, upper_green, upper_blue, lower_red, lower_green, lower_blue)
		}
		fmt.Println("\x1b[0m")
	}
}

func main() {
	flag.Parse()
	filenames := flag.Args()

	for i, filename := range filenames {
		if i > 0 {
			fmt.Println()
		}

		fmt.Printf("%s:\n", filename)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		source_image, _, err := image.Decode(file)
		file.Close()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		PrintImage(source_image)
	}
}
