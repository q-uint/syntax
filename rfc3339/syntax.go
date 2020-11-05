// This file is generated - do not edit.

package rfc3339

import (
	"github.com/elimity-com/abnf/core"

	"github.com/elimity-com/abnf/operators"
)

// date-fullyear = 4DIGIT
func DateFullyear(s []byte) operators.Alternatives {
	return operators.RepeatN("date-fullyear", '\x04', core.DIGIT())(s)
}

// date-mday = 2DIGIT
func DateMday(s []byte) operators.Alternatives {
	return operators.RepeatN("date-mday", '\x02', core.DIGIT())(s)
}

// date-month = 2DIGIT
func DateMonth(s []byte) operators.Alternatives {
	return operators.RepeatN("date-month", '\x02', core.DIGIT())(s)
}

// date-time = full-date "T" full-time
func DateTime(s []byte) operators.Alternatives {
	return operators.Concat(
		"date-time",
		FullDate,
		operators.String("T", "T"),
		FullTime,
	)(s)
}

// full-date = date-fullyear "-" date-month "-" date-mday
func FullDate(s []byte) operators.Alternatives {
	return operators.Concat(
		"full-date",
		DateFullyear,
		operators.String("-", "-"),
		DateMonth,
		operators.String("-", "-"),
		DateMday,
	)(s)
}

// full-time = partial-time time-offset
func FullTime(s []byte) operators.Alternatives {
	return operators.Concat(
		"full-time",
		PartialTime,
		TimeOffset,
	)(s)
}

// partial-time = time-hour ":" time-minute ":" time-second [time-secfrac]
func PartialTime(s []byte) operators.Alternatives {
	return operators.Concat(
		"partial-time",
		TimeHour,
		operators.String(":", ":"),
		TimeMinute,
		operators.String(":", ":"),
		TimeSecond,
		operators.Optional("[time-secfrac]", TimeSecfrac),
	)(s)
}

// time-hour = 2DIGIT
func TimeHour(s []byte) operators.Alternatives {
	return operators.RepeatN("time-hour", '\x02', core.DIGIT())(s)
}

// time-minute = 2DIGIT
func TimeMinute(s []byte) operators.Alternatives {
	return operators.RepeatN("time-minute", '\x02', core.DIGIT())(s)
}

// time-numoffset = ("+" / "-") time-hour ":" time-minute
func TimeNumoffset(s []byte) operators.Alternatives {
	return operators.Concat(
		"time-numoffset",
		operators.Alts(
			"\"+\" / \"-\"",
			operators.String("+", "+"),
			operators.String("-", "-"),
		),
		TimeHour,
		operators.String(":", ":"),
		TimeMinute,
	)(s)
}

// time-offset = "Z" / time-numoffset
func TimeOffset(s []byte) operators.Alternatives {
	return operators.Alts(
		"time-offset",
		operators.String("Z", "Z"),
		TimeNumoffset,
	)(s)
}

// time-secfrac = "." 1*DIGIT
func TimeSecfrac(s []byte) operators.Alternatives {
	return operators.Concat(
		"time-secfrac",
		operators.String(".", "."),
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
	)(s)
}

// time-second = 2DIGIT
func TimeSecond(s []byte) operators.Alternatives {
	return operators.RepeatN("time-second", '\x02', core.DIGIT())(s)
}
