# Go-ImageQuant
[![Go Report Card](https://goreportcard.com/badge/github.com/bbiao/go-imagequant)](https://goreportcard.com/report/github.com/bbiao/go-imagequant) [![Build Status](https://travis-ci.org/bbiao/go-imagequant.svg?branch=master)](https://travis-ci.org/bbiao/go-imagequant) [![GoDoc](https://godoc.org/github.com/bbiao/go-imagequant?status.svg)](https://godoc.org/github.com/bbiao/go-imagequant)
## ABOUT
This is Go bindings for libimagequant.

Libimagequant is the backend of pngquant app. It provides a high level of png image compression.

Libimagequant is a library for lossy recompression of PNG images to reduce their filesize.  This go-imagequant project is a set of bindings for libimagequant to enable its use from the Go programming language.

This binding was written by hand. The result is somewhat more idiomatic than an automated conversion, but some  defer foo.Release() calls are required for memory management.

This project forked from [https://github.com/ultimate-guitar/go-imagequant](https://github.com/ultimate-guitar/go-imagequant) repo.

This project forked from [https://code.ivysaur.me/go-imagequant/](https://code.ivysaur.me/go-imagequant/) repo.

## USAGE
Usage example is provided by a sample utility cmd/gopngquant which mimics some functionality of the upstream pngquant.

The sample utility has the following options:

```
Usage of gopngquant:
  -colors int
        number of colors (default 256)
  -compression int
        compression level (DefaultCompression = 0, NoCompression = -1, BestSpeed = -2, BestCompression = -3) (default -3)
  -ext string
        file name suffix append to input file name if output not presents
  -input string
        input filename
  -nofs
        disable Floyd-Steinberg dithering
  -output string
        output filename
  -quality string
        don't save below min, use fewer colors below max (0-100) (default "0-100")
  -speed int
        speed (1 slowest, 10 fastest) (default 3)
  -y    overwrite existing output file
  -version
```

## BUILDING

Install libimagequant, in ubuntu, just run `sudo apt-get install libimagequant-dev`

Install a C11 compiler and simply `go get github.com/bbiao/go-imagequant`.

## LICENSE
I am releasing this binding under the ISC license, however, libimagequant itself is released under GPLv3-or-later and/or commercial licenses. You must comply with the terms of such a license when using this binding in a Go project.

## CHANGELOG

See [CHANGELOG](CHANGELOG.md).
