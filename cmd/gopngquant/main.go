/*
Copyright (c) 2016, The go-imagequant author(s)

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND ISC DISCLAIMS ALL WARRANTIES WITH REGARD TO
THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS.
IN NO EVENT SHALL ISC BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR
CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA
OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS
ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS
SOFTWARE.
*/

package main

import (
	"flag"
	"fmt"
	"github.com/bbiao/go-imagequant"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	shouldDisplayVersion bool
	input                string
	output               string
	speed                int
	compression          int
	quality              string
	colors               int
	nofs                 bool
	force                bool
	fileSuffix           string
)

func init() {
	flag.BoolVar(&shouldDisplayVersion, "version", false, "")
	flag.StringVar(&input, "input", "", "input filename")
	flag.StringVar(&output, "output", "", "output filename")
	flag.IntVar(&speed, "speed", 3, "speed (1 slowest, 10 fastest)")
	flag.IntVar(&compression, "compression", -3, "compression level (DefaultCompression = 0, NoCompression = -1, BestSpeed = -2, BestCompression = -3)")
	flag.StringVar(&quality, "quality", "0-100", "don't save below min, use fewer colors below max (0-100)")
	flag.IntVar(&colors, "colors", 256, "number of colors")
	flag.BoolVar(&nofs, "nofs", false, "disable Floyd-Steinberg dithering")
	flag.BoolVar(&force, "y", false, "overwrite existing output file")
	flag.StringVar(&fileSuffix, "ext", "", "file name suffix append to input file name if output not presents")

	flag.Parse()

	if shouldDisplayVersion {
		fmt.Printf("libimagequant '%s' (%d)\n", imagequant.GetLibraryVersionString(), imagequant.GetLibraryVersion())
		os.Exit(0)
	}
}

func crushFile(sourcefile, destfile string, opts *imagequant.Options) error {
	sourceFh, err := os.OpenFile(sourcefile, os.O_RDONLY, 0444)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %s", err.Error())
	}
	defer sourceFh.Close()

	image, err := ioutil.ReadAll(sourceFh)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %s", err.Error())
	}

	optiImage, err := imagequant.Crush(image, opts)
	if err != nil {
		return fmt.Errorf("imagequant.Crush: %s", err.Error())
	}

	destFh, err := os.OpenFile(destfile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %s", err.Error())
	}
	defer destFh.Close()

	destFh.Write(optiImage)
	return nil
}

func parseCompressionLevel() png.CompressionLevel {
	var cLevel png.CompressionLevel
	switch compression {
	case 0:
		cLevel = png.DefaultCompression
	case -1:
		cLevel = png.NoCompression
	case -2:
		cLevel = png.BestSpeed
	case -3:
		cLevel = png.BestCompression
	default:
		cLevel = png.BestCompression
	}
	return cLevel
}

func parseQuality() (int, int, error) {
	if quality == "" {
		return 0, 100, nil
	}

	qpair := strings.Split(quality, "-")
	if len(qpair) != 2 {
		return 0, 0, fmt.Errorf("format error: %s", quality)
	}
	minQuality, err := strconv.Atoi(qpair[0])
	if err != nil {
		return 0, 0, fmt.Errorf("min quality error: %s", qpair[0])
	}
	maxQuality, err := strconv.Atoi(qpair[1])
	if err != nil {
		return 0, 0, fmt.Errorf("max quality error: %s", qpair[1])
	}
	return minQuality, maxQuality, nil
}

func main() {
	opts := imagequant.DefaultOptions
	opts.Colors = colors
	opts.Speed = speed
	opts.Compression = parseCompressionLevel()
	minQuality, maxQuality, err := parseQuality()
	if err != nil {
		fmt.Printf("quality parse error, %s\n", err.Error())
		os.Exit(-1)
	}
	opts.MinQuality, opts.MaxQuality = minQuality, maxQuality
	if nofs {
		opts.Dithering = 0.
	}

	if input == "" {
		fmt.Println("must specified input file")
		os.Exit(-1)
	}

	if output == ""  {
		if fileSuffix == "" {
			fmt.Println("must specified output or ext")
			os.Exit(-1)
		}

		pos := strings.LastIndex(input, ".")
		if pos == -1 {
			output = input + fileSuffix
		} else {
			output = input[0:pos] + fileSuffix + input[pos:]
		}
	}

	if !force {
		if _, err := os.Stat(output); err == nil {
			fmt.Printf("output file %s already exists\n", output)
			os.Exit(-1)
		}
	}

	err = crushFile(input, output, opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
