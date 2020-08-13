// This file is generated - do not edit.

package iso8601

import "github.com/elimity-com/abnf/operators"

// period = period-explicit / period-start / period-end
func Period(s []byte) operators.Alternatives {
	return operators.Alts(
		"period",
		PeriodExplicit,
		PeriodStart,
		PeriodEnd,
	)(s)
}

// period-end = duration "/" iso-date-time
func PeriodEnd(s []byte) operators.Alternatives {
	return operators.Concat(
		"period-end",
		Duration,
		operators.String("/", "/"),
		IsoDateTime,
	)(s)
}

// period-explicit = iso-date-time "/" iso-date-time
func PeriodExplicit(s []byte) operators.Alternatives {
	return operators.Concat(
		"period-explicit",
		IsoDateTime,
		operators.String("/", "/"),
		IsoDateTime,
	)(s)
}

// period-start = iso-date-time "/" duration
func PeriodStart(s []byte) operators.Alternatives {
	return operators.Concat(
		"period-start",
		IsoDateTime,
		operators.String("/", "/"),
		Duration,
	)(s)
}
