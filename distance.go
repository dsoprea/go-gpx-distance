package gpxdistance

import (
	"io"
	"sort"

	"github.com/asmarques/geodist"

	"github.com/dsoprea/go-gpx"
	"github.com/dsoprea/go-gpx/reader"
	"github.com/dsoprea/go-logging/v2"
)

// sortableTrackPoints allows us to sort the points by timestamp
type sortableTrackPoints []gpxcommon.TrackPoint

func (stp sortableTrackPoints) Len() int {
	return len(stp)
}

func (stp sortableTrackPoints) Swap(i, j int) {
	stp[i], stp[j] = stp[j], stp[i]
}

func (stp sortableTrackPoints) Less(i, j int) bool {
	return stp[i].Time.Before(stp[j].Time)
}

// Calculate returns the number of kilometers traveled GPX data given a reader
// for that data. It first sorts by timestamp.
func Calculate(r io.Reader) (totalDistanceKm float64, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	points, err := gpxreader.ExtractTrackPoints(r)
	log.PanicIf(err)

	sort.Sort(sortableTrackPoints(points))

	var zeroPoint gpxcommon.TrackPoint
	var lastPoint gpxcommon.TrackPoint
	for _, currentPoint := range points {

		if lastPoint != zeroPoint {
			p1 := geodist.Point{
				Lat:  lastPoint.LatitudeDecimal,
				Long: lastPoint.LongitudeDecimal,
			}

			p2 := geodist.Point{
				Lat:  currentPoint.LatitudeDecimal,
				Long: currentPoint.LongitudeDecimal,
			}

			// Calculate distance in kilometers
			d := geodist.HaversineDistance(p1, p2)
			totalDistanceKm += d
		}

		lastPoint = currentPoint
	}

	return totalDistanceKm, nil
}
