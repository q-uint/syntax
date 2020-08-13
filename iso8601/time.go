// This file is generated - do not edit.

package iso8601

import (
	core "github.com/elimity-com/abnf/core"
	"github.com/elimity-com/abnf/operators"
)

// iso-date-time = date "T" time
func IsoDateTime(s []byte) operators.Alternatives {
	return operators.Concat(
		"iso-date-time",
		Date,
		operators.String("T", "T"),
		Time,
	)(s)
}

// time = timespec-base [time-fraction] [time-zone]
func Time(s []byte) operators.Alternatives {
	return operators.Concat(
		"time",
		TimespecBase,
		operators.Optional("[time-fraction]", TimeFraction),
		operators.Optional("[time-zone]", TimeZone),
	)(s)
}

// time-fraction = ("," / ".") 1*DIGIT
func TimeFraction(s []byte) operators.Alternatives {
	return operators.Concat(
		"time-fraction",
		operators.Alts(
			"\",\" / \".\"",
			operators.String(",", ","),
			operators.String(".", "."),
		),
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
	)(s)
}

// time-hour = 2DIGIT
func TimeHour(s []byte) operators.Alternatives {
	return operators.RepeatN("time-hour", 2, core.DIGIT())(s)
}

// time-minute = 2DIGIT
func TimeMinute(s []byte) operators.Alternatives {
	return operators.RepeatN("time-minute", 2, core.DIGIT())(s)
}

// time-numoffset = ("+" / "-") time-hour [[":"] time-minute]
func TimeNumoffset(s []byte) operators.Alternatives {
	return operators.Concat(
		"time-numoffset",
		operators.Alts(
			"\"+\" / \"-\"",
			operators.String("+", "+"),
			operators.String("-", "-"),
		),
		TimeHour,
		operators.Optional("[[\":\"] time-minute]", operators.Concat(
			"[\":\"] time-minute",
			operators.Optional("[\":\"]", operators.String(":", ":")),
			TimeMinute,
		)),
	)(s)
}

// time-second = 2DIGIT
func TimeSecond(s []byte) operators.Alternatives {
	return operators.RepeatN("time-second", 2, core.DIGIT())(s)
}

// time-zone = "Z" / time-numoffset
func TimeZone(s []byte) operators.Alternatives {
	return operators.Alts(
		"time-zone",
		operators.String("Z", "Z"),
		TimeNumoffset,
	)(s)
}

// timeopt-hour = "-" / (time-hour [":"])
func TimeoptHour(s []byte) operators.Alternatives {
	return operators.Alts(
		"timeopt-hour",
		operators.String("-", "-"),
		operators.Concat(
			"time-hour [\":\"]",
			TimeHour,
			operators.Optional("[\":\"]", operators.String(":", ":")),
		),
	)(s)
}

// timeopt-minute = "-" / (time-minute [":"])
func TimeoptMinute(s []byte) operators.Alternatives {
	return operators.Alts(
		"timeopt-minute",
		operators.String("-", "-"),
		operators.Concat(
			"time-minute [\":\"]",
			TimeMinute,
			operators.Optional("[\":\"]", operators.String(":", ":")),
		),
	)(s)
}

// timespec-base = timespec-hour / timespec-minute / timespec-second
func TimespecBase(s []byte) operators.Alternatives {
	return operators.Alts(
		"timespec-base",
		TimespecHour,
		TimespecMinute,
		TimespecSecond,
	)(s)
}

// timespec-hour = time-hour [[":"] time-minute [[":"] time-second]]
func TimespecHour(s []byte) operators.Alternatives {
	return operators.Concat(
		"timespec-hour",
		TimeHour,
		operators.Optional("[[\":\"] time-minute [[\":\"] time-second]]", operators.Concat(
			"[\":\"] time-minute [[\":\"] time-second]",
			operators.Optional("[\":\"]", operators.String(":", ":")),
			TimeMinute,
			operators.Optional("[[\":\"] time-second]", operators.Concat(
				"[\":\"] time-second",
				operators.Optional("[\":\"]", operators.String(":", ":")),
				TimeSecond,
			)),
		)),
	)(s)
}

// timespec-minute = timeopt-hour time-minute [[":"] time-second]
func TimespecMinute(s []byte) operators.Alternatives {
	return operators.Concat(
		"timespec-minute",
		TimeoptHour,
		TimeMinute,
		operators.Optional("[[\":\"] time-second]", operators.Concat(
			"[\":\"] time-second",
			operators.Optional("[\":\"]", operators.String(":", ":")),
			TimeSecond,
		)),
	)(s)
}

// timespec-second = "-" timeopt-minute time-second
func TimespecSecond(s []byte) operators.Alternatives {
	return operators.Concat(
		"timespec-second",
		operators.String("-", "-"),
		TimeoptMinute,
		TimeSecond,
	)(s)
}
