package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/dsoprea/go-logging/v2"

	"github.com/dsoprea/go-gpx-distance"
)

const (
	kilometersPerMile = 1.60934

	gpxExtension = ".gpx"
)

var (
	mainLogger = log.NewLogger("main.main")
)

type parameters struct {
	Path      string `short:"p" long:"path" required:"true" description:"Path of GPX files. Will scan recursively."`
	IsVerbose bool   `short:"v" long:"verbose" description:"Print logging"`
}

var (
	arguments = new(parameters)
)

func main() {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err := errRaw.(error)
			log.PrintError(err)

			os.Exit(-2)
		}
	}()

	_, err := flags.Parse(arguments)
	if err != nil {
		os.Exit(-1)
	}

	if arguments.IsVerbose == true {
		cla := log.NewConsoleLogAdapter()
		log.AddAdapter("console", cla)

		scp := log.NewStaticConfigurationProvider()
		scp.SetLevelName("debug")

		log.LoadConfiguration(scp)
	}

	var totalDistanceKm float64

	cb := func(filepath_ string, de fs.DirEntry, err error) (err2 error) {
		defer func() {
			if state := recover(); state != nil {
				err2 = log.Wrap(state.(error))
			}
		}()

		log.PanicIf(err)

		// Filter for GPX files

		filename := path.Base(filepath_)

		extension := filepath.Ext(filename)
		extension = strings.ToLower(extension)

		if extension != gpxExtension {
			return nil
		}

		// Open

		f, err := os.Open(filepath_)
		log.PanicIf(err)

		defer f.Close()

		// Calculate

		distanceKm, err := gpxdistance.Calculate(f)
		log.PanicIf(err)

		// Print and add

		fmt.Fprintf(os.Stderr, "%s: %.1f\n",
			filepath_, distanceKm)

		totalDistanceKm += distanceKm

		return nil
	}

	err = filepath.WalkDir(arguments.Path, fs.WalkDirFunc(cb))
	log.PanicIf(err)

	fmt.Fprintf(os.Stderr, "\n")

	// Print totals

	totalDistanceM := totalDistanceKm / kilometersPerMile

	fmt.Printf("%f kilometers\n", totalDistanceKm)
	fmt.Printf("%f miles\n", totalDistanceM)
	fmt.Printf("\n")
}
