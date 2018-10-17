package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"io"
	"os"

	"flag"
	"image/png"
)

func main() {

	var (
		format = flag.String("format", "jpg", "output image format")
	)
	flag.Parse()
	if err := toJPEG(os.Stdin, os.Stdout, format); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer, format *string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch *format {
	case "jpeg":
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	}
	return fmt.Errorf("unsupported format %s", *format)
}
