package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func convertAndResize(filePath string, width uint, height uint) error {
	// Read the WebP file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Decode the WebP image
	img, err := webp.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	// Resize the image
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Save the resized image as PNG
	outFile, err := os.Create(strings.TrimSuffix(filePath, ".webp") + ".png")
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, resizedImg)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./app <directory-path> <width>x<height>")
		return
	}

	dirPath := os.Args[1]
	size := strings.Split(os.Args[2], "x")
	if len(size) != 2 {
		fmt.Println("Invalid size format. It should be like 32x32.")
		return
	}

	width, err := strconv.ParseUint(size[0], 10, 32)
	if err != nil {
		fmt.Println("Invalid width:", err)
		return
	}

	height, err := strconv.ParseUint(size[1], 10, 32)
	if err != nil {
		fmt.Println("Invalid height:", err)
		return
	}

	// Iterate over files in the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".webp" {
			filePath := filepath.Join(dirPath, f.Name())
			err := convertAndResize(filePath, uint(width), uint(height))
			if err != nil {
				fmt.Printf("Failed to convert %s: %v\n", f.Name(), err)
			} else {
				fmt.Printf("Converted %s successfully\n", f.Name())
			}
		}
	}
}
