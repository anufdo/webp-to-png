# WebP to PNG Converter with Resizing

This CLI application is written in Go. It's designed to batch convert `.webp` images to `.png` format and simultaneously resize them based on user-specified dimensions.

## Features

- Batch converts all `.webp` images in a specified directory.
- Resizes images during the conversion process.
- Command-line interface for easy use.

## Prerequisites

- Go (at least version 1.11, for module support)

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/anufdo/webp-to-png.git
   cd webp-to-png-converter
   ``````

2. Install the required packages:
    ```sh
    go mod tidy 
    ```

3. Build the application:
    ```sh
    go build converter.go
    ```

## Usage

To convert .webp images in a directory and resize them:
    
    ./converter /path/to/directory WIDTHxHEIGHT

Replace /path/to/directory with the directory containing your .webp images, and WIDTHxHEIGHT with the desired dimensions, e.g., 32x32.

## License
MIT 

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you'd like to change.

## Acknowledgements

- [webp by chai2010](https://github.com/chai2010/webp) for WebP decoding.
- [resize by nfnt](https://github.com/nfnt/resize) for image resizing.


    

