package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/dsoprea/go-logging/v2"

	"github.com/dsoprea/go-gpx-distance"
)

const (
	kilometersPerMile = 1.60934
)

var (
	mainLogger = log.NewLogger("main.main")
)

type parameters struct {
	Filepath  string `short:"f" long:"filepath" required:"true" description:"File-path of GPX file"`
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

	f, err := os.Open(arguments.Filepath)
	log.PanicIf(err)

	defer f.Close()

	totalDistanceKm, err := gpxdistance.Calculate(f)
	log.PanicIf(err)

	totalDistanceM := totalDistanceKm / kilometersPerMile

	fmt.Printf("%f kilometers\n", totalDistanceKm)
	fmt.Printf("%f miles\n", totalDistanceM)
	fmt.Printf("\n")
}
