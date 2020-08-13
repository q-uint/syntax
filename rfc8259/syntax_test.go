package syntax

import (
	"testing"
)

func TestExamples(t *testing.T) {
	// https://tools.ietf.org/html/rfc8259#section-13
	for _, test := range [][]byte{
		[]byte(`{
        "Image": {
            "Width":  800,
            "Height": 600,
            "Title":  "View from 15th Floor",
            "Thumbnail": {
                "Url":    "http://www.example.com/image/481989943",
                "Height": 125,
                "Width":  100
            },
            "Animated" : false,
            "IDs": [116, 943, 234, 38793]
		}
	}`),
		[]byte(`[
        {
           "precision": "zip",
           "Latitude":  37.7668,
           "Longitude": -122.3959,
           "Address":   "",
           "City":      "SAN FRANCISCO",
           "State":     "CA",
           "Zip":       "94107",
           "Country":   "US"
        },
        {
           "precision": "zip",
           "Latitude":  37.371991,
           "Longitude": -122.026020,
           "Address":   "",
           "City":      "SUNNYVALE",
           "State":     "CA",
           "Zip":       "94085",
           "Country":   "US"
        }
	]`),
		[]byte(`"Hello world!"`),
		[]byte(`42`),
		[]byte(`true`),
	} {
		if JSONText(test).Best() == nil {
			t.Error("could not parse JSON")
		}
	}

}
