// This file is generated - do not edit.

package iso8601

import (
	core "github.com/elimity-com/abnf/core"
	"github.com/elimity-com/abnf/operators"
)

// date = datespec-full / datespec-year / datespec-month / datespec-mday / datespec-week / datespec-wday / datespec-yday
func Date(s []byte) operators.Alternatives {
	return operators.Alts(
		"date",
		DatespecFull,
		DatespecYear,
		DatespecMonth,
		DatespecMday,
		DatespecWeek,
		DatespecWday,
		DatespecYday,
	)(s)
}

// date-century = 2DIGIT
func DateCentury(s []byte) operators.Alternatives {
	return operators.RepeatN("date-century", 2, core.DIGIT())(s)
}

// date-decade = DIGIT
func DateDecade(s []byte) operators.Alternatives {
	return core.DIGIT()(s)
}

// date-fullyear = date-century date-year
func DateFullyear(s []byte) operators.Alternatives {
	return operators.Concat(
		"date-fullyear",
		DateCentury,
		DateYear,
	)(s)
}

// date-mday = 2DIGIT
func DateMday(s []byte) operators.Alternatives {
	return operators.RepeatN("date-mday", 2, core.DIGIT())(s)
}

// date-month = 2DIGIT
func DateMonth(s []byte) operators.Alternatives {
	return operators.RepeatN("date-month", 2, core.DIGIT())(s)
}

// date-subdecade = DIGIT
func DateSubdecade(s []byte) operators.Alternatives {
	return core.DIGIT()(s)
}

// date-wday = DIGIT
func DateWday(s []byte) operators.Alternatives {
	return core.DIGIT()(s)
}

// date-week = 2DIGIT
func DateWeek(s []byte) operators.Alternatives {
	return operators.RepeatN("date-week", 2, core.DIGIT())(s)
}

// date-yday = 3DIGIT
func DateYday(s []byte) operators.Alternatives {
	return operators.RepeatN("date-yday", 3, core.DIGIT())(s)
}

// date-year = date-decade date-subdecade
func DateYear(s []byte) operators.Alternatives {
	return operators.Concat(
		"date-year",
		DateDecade,
		DateSubdecade,
	)(s)
}

// dateopt-century = "-" / date-century
func DateoptCentury(s []byte) operators.Alternatives {
	return operators.Alts(
		"dateopt-century",
		operators.String("-", "-"),
		DateCentury,
	)(s)
}

// dateopt-fullyear = "-" / datepart-fullyear
func DateoptFullyear(s []byte) operators.Alternatives {
	return operators.Alts(
		"dateopt-fullyear",
		operators.String("-", "-"),
		DatepartFullyear,
	)(s)
}

// dateopt-month = "-" / (date-month ["-"])
func DateoptMonth(s []byte) operators.Alternatives {
	return operators.Alts(
		"dateopt-month",
		operators.String("-", "-"),
		operators.Concat(
			"date-month [\"-\"]",
			DateMonth,
			operators.Optional("[\"-\"]", operators.String("-", "-")),
		),
	)(s)
}

// dateopt-week = "-" / (date-week ["-"])
func DateoptWeek(s []byte) operators.Alternatives {
	return operators.Alts(
		"dateopt-week",
		operators.String("-", "-"),
		operators.Concat(
			"date-week [\"-\"]",
			DateWeek,
			operators.Optional("[\"-\"]", operators.String("-", "-")),
		),
	)(s)
}

// dateopt-year = "-" / (date-year ["-"])
func DateoptYear(s []byte) operators.Alternatives {
	return operators.Alts(
		"dateopt-year",
		operators.String("-", "-"),
		operators.Concat(
			"date-year [\"-\"]",
			DateYear,
			operators.Optional("[\"-\"]", operators.String("-", "-")),
		),
	)(s)
}

// datepart-fullyear = [date-century] date-year ["-"]
func DatepartFullyear(s []byte) operators.Alternatives {
	return operators.Concat(
		"datepart-fullyear",
		operators.Optional("[date-century]", DateCentury),
		DateYear,
		operators.Optional("[\"-\"]", operators.String("-", "-")),
	)(s)
}

// datepart-ptyear = "-" [date-subdecade ["-"]]
func DatepartPtyear(s []byte) operators.Alternatives {
	return operators.Concat(
		"datepart-ptyear",
		operators.String("-", "-"),
		operators.Optional("[date-subdecade [\"-\"]]", operators.Concat(
			"date-subdecade [\"-\"]",
			DateSubdecade,
			operators.Optional("[\"-\"]", operators.String("-", "-")),
		)),
	)(s)
}

// datepart-wkyear = datepart-ptyear / datepart-fullyear
func DatepartWkyear(s []byte) operators.Alternatives {
	return operators.Alts(
		"datepart-wkyear",
		DatepartPtyear,
		DatepartFullyear,
	)(s)
}

// datespec-full = datepart-fullyear date-month ["-"] date-mday
func DatespecFull(s []byte) operators.Alternatives {
	return operators.Concat(
		"datespec-full",
		DatepartFullyear,
		DateMonth,
		operators.Optional("[\"-\"]", operators.String("-", "-")),
		DateMday,
	)(s)
}

// datespec-mday = "--" dateopt-month date-mday
func DatespecMday(s []byte) operators.Alternatives {
	return operators.Concat(
		"datespec-mday",
		operators.String("--", "--"),
		DateoptMonth,
		DateMday,
	)(s)
}

// datespec-month = "-" dateopt-year date-month [["-"] date-mday]
func DatespecMonth(s []byte) operators.Alternatives {
	return operators.Concat(
		"datespec-month",
		operators.String("-", "-"),
		DateoptYear,
		DateMonth,
		operators.Optional("[[\"-\"] date-mday]", operators.Concat(
			"[\"-\"] date-mday",
			operators.Optional("[\"-\"]", operators.String("-", "-")),
			DateMday,
		)),
	)(s)
}

// datespec-wday = "---" date-wday
func DatespecWday(s []byte) operators.Alternatives {
	return operators.Concat(
		"datespec-wday",
		operators.String("---", "---"),
		DateWday,
	)(s)
}

// datespec-week = datepart-wkyear "W" (date-week / dateopt-week date-wday)
func DatespecWeek(s []byte) operators.Alternatives {
	return operators.Concat(
		"datespec-week",
		DatepartWkyear,
		operators.String("W", "W"),
		operators.Alts(
			"date-week / dateopt-week date-wday",
			DateWeek,
			operators.Concat(
				"dateopt-week date-wday",
				DateoptWeek,
				DateWday,
			),
		),
	)(s)
}

// datespec-yday = dateopt-fullyear date-yday
func DatespecYday(s []byte) operators.Alternatives {
	return operators.Concat(
		"datespec-yday",
		DateoptFullyear,
		DateYday,
	)(s)
}

// datespec-year = date-century / dateopt-century date-year
func DatespecYear(s []byte) operators.Alternatives {
	return operators.Alts(
		"datespec-year",
		DateCentury,
		operators.Concat(
			"dateopt-century date-year",
			DateoptCentury,
			DateYear,
		),
	)(s)
}
