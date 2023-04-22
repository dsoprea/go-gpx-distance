[![Go Report Card](https://goreportcard.com/badge/github.com/dsoprea/go-gpx-distance)](https://goreportcard.com/report/github.com/dsoprea/go-gpx-distance)
[![GoDoc](https://godoc.org/github.com/dsoprea/go-gpx-distance?status.svg)](https://godoc.org/github.com/dsoprea/go-gpx-distance)

# Overview

This provides a library and tools to count the distance traveled in GPX data. The library provides a function that takes an `io.Reader`. Alternatively, the "go-gpx-distance" command takes a single GPX file and "go-gpx-distance-path" takes a path to traverse recursively.

# Example

Given a `io.Reader`, calculate the distance traveled:

```
distanceKm, err := Calculate(r)
log.PanicIf(err)
```

# Installing Commands

If you just want the commands and not the sourcecode, install via:

```
$ go install github.com/dsoprea/go-gpx-distance/command/go-gpx-distance@latest
$ go install github.com/dsoprea/go-gpx-distance/command/go-gpx-distance-path@latest
```

# XML Encoding

Note that, since Go doesn't supporting XML decoding for any encoding but UTF-8, any GPX files with any other encoding but this will yield empty data. If you have a sufficiently similar encoding, such as ISO-8859-1, the current, best solution would just be to change the encoding to "UTF-8":

From:
```
<?xml version="1.0" encoding="ISO-8859-1" ?>
```

To:
```
<?xml version="1.0" encoding="UTF-8" ?>
```
