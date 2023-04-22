package gpxdistance

import (
	"bytes"
	"math"
	"testing"

	"github.com/dsoprea/go-logging/v2"
)

const (
	testData = `
<?xml version="1.0"
standalone="yes"?>
<?xml-stylesheet type="text/xsl" href="details.xsl"?>
<gpx
 version="1.0"
 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
 xmlns="http://www.topografix.com/GPX/1/0"
 xmlns:topografix="http://www.topografix.com/GPX/Private/TopoGrafix/0/1" xsi:schemaLocation="http://www.topografix.com/GPX/1/0 http://www.topografix.com/GPX/1/0/gpx.xsd http://www.topografix.com/GPX/Private/TopoGrafix/0/1 http://www.topografix.com/GPX/Private/TopoGrafix/0/1/topografix.xsd">
<trk>
<name><![CDATA[2023-04-20 19:45]]></name>
<desc><![CDATA[]]></desc>
<number>2206</number>
<topografix:color>c0c0c0</topografix:color>
<trkseg>
<trkpt lat="41.266009" lon="28.734982">
<ele>140.987060546875</ele>
<time>2023-04-20T16:45:25Z</time>
</trkpt>
<trkpt lat="41.008917" lon="28.977601">
<ele>83.51824951171875</ele>
<time>2023-04-20T17:42:26Z</time>
</trkpt>
<trkpt lat="41.008895" lon="28.977629">
<ele>82.16912841796875</ele>
<time>2023-04-20T17:42:36Z</time>
</trkpt>
<trkpt lat="41.008919" lon="28.977627">
<ele>79.72216796875</ele>
<time>2023-04-20T17:43:06Z</time>
</trkpt>
<trkpt lat="41.008906" lon="28.977645">
<ele>79.404296875</ele>
<time>2023-04-20T17:43:16Z</time>
</trkpt>
<trkpt lat="41.008892" lon="28.977661">
<ele>78.26611328125</ele>
<time>2023-04-20T17:43:26Z</time>
</trkpt>
</trkseg>
</trk>
</gpx>
`
)

func TestDistance(t *testing.T) {

	b := bytes.NewBufferString(testData)

	totalDistanceKm, err := Calculate(b)
	log.PanicIf(err)

	expected := 35.0823539754785

	if totalDistanceKm < expected || totalDistanceKm > math.Nextafter(expected, expected+1) {
		t.Fatalf("Distance not right: %f != %f", totalDistanceKm, expected)
	}
}
