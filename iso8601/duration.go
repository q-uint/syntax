// This file is generated - do not edit.

package iso8601

import (
	"github.com/elimity-com/abnf/core"

	"github.com/elimity-com/abnf/operators"
)

// dur-date = (dur-day / dur-month / dur-year) [dur-time]
func DurDate(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-date",
		operators.Alts(
			"dur-day / dur-month / dur-year",
			DurDay,
			DurMonth,
			DurYear,
		),
		operators.Optional("[dur-time]", DurTime),
	)(s)
}

// dur-day = 1*DIGIT "D"
func DurDay(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-day",
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
		operators.String("D", "D"),
	)(s)
}

// dur-hour = 1*DIGIT "H" [dur-minute]
func DurHour(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-hour",
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
		operators.String("H", "H"),
		operators.Optional("[dur-minute]", DurMinute),
	)(s)
}

// dur-minute = 1*DIGIT "M" [dur-second]
func DurMinute(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-minute",
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
		operators.String("M", "M"),
		operators.Optional("[dur-second]", DurSecond),
	)(s)
}

// dur-month = 1*DIGIT "M" [dur-day]
func DurMonth(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-month",
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
		operators.String("M", "M"),
		operators.Optional("[dur-day]", DurDay),
	)(s)
}

// dur-second = 1*DIGIT "S"
func DurSecond(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-second",
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
		operators.String("S", "S"),
	)(s)
}

// dur-time = "T" (dur-hour / dur-minute / dur-second)
func DurTime(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-time",
		operators.String("T", "T"),
		operators.Alts(
			"dur-hour / dur-minute / dur-second",
			DurHour,
			DurMinute,
			DurSecond,
		),
	)(s)
}

// dur-week = 1*DIGIT "W"
func DurWeek(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-week",
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
		operators.String("W", "W"),
	)(s)
}

// dur-year = 1*DIGIT "Y" [dur-month]
func DurYear(s []byte) operators.Alternatives {
	return operators.Concat(
		"dur-year",
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
		operators.String("Y", "Y"),
		operators.Optional("[dur-month]", DurMonth),
	)(s)
}

// duration = "P" (dur-date / dur-time / dur-week)
func Duration(s []byte) operators.Alternatives {
	return operators.Concat(
		"duration",
		operators.String("P", "P"),
		operators.Alts(
			"dur-date / dur-time / dur-week",
			DurDate,
			DurTime,
			DurWeek,
		),
	)(s)
}
