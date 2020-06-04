package syntax

import . "github.com/elimity-com/abnf/operators"

var (
	// Common Elements
	// https://tools.ietf.org/html/rfc2396#section-1.6
	alpha    = Alts(`alpha`, lowalpha, upalpha)
	lowalpha = Range(`lowalpha`, 'a', 'z')
	upalpha  = Range(`upalpha`, 'A', 'Z')
	digit    = Range(`digit`, '0', '9')
	alphanum = Alts(`alphanum`, alpha, digit)
)
