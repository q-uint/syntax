package syntax

import "testing"

func TestExamples(t *testing.T) {
	// https://tools.ietf.org/html/rfc3339#section-5.8
	for _, test := range [][]byte{
		[]byte("1985-04-12T23:20:50.52Z"),
		[]byte("1996-12-19T16:39:57-08:00"),
		[]byte("1990-12-31T23:59:60Z"),
		[]byte("1990-12-31T15:59:60-08:00"),
		[]byte("1937-01-01T12:00:27.87+00:20"),
	} {
		if DateTime(test).Best() == nil {
			t.Error("could not parse DateTime")
		}
	}
}
