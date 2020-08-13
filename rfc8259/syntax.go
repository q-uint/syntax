// This file is generated - do not edit.

package rfc8259

import (
	core "github.com/elimity-com/abnf/core"
	"github.com/elimity-com/abnf/operators"
)

// JSON-text = ws value ws
func JSONText(s []byte) operators.Alternatives {
	return operators.Concat(
		"JSON-text",
		Ws,
		Value,
		Ws,
	)(s)
}

// array = begin-array [ value *( value-separator value ) ] end-array
func Array(s []byte) operators.Alternatives {
	return operators.Concat(
		"array",
		BeginArray,
		operators.Optional("[ value *( value-separator value ) ]", operators.Concat(
			"value *( value-separator value )",
			Value,
			operators.Repeat0Inf("*( value-separator value )", operators.Concat(
				"value-separator value",
				ValueSeparator,
				Value,
			)),
		)),
		EndArray,
	)(s)
}

// begin-array = ws %x5B ws
func BeginArray(s []byte) operators.Alternatives {
	return operators.Concat(
		"begin-array",
		Ws,
		operators.Terminal("%x5B", []byte{91}),
		Ws,
	)(s)
}

// begin-object = ws %x7B ws
func BeginObject(s []byte) operators.Alternatives {
	return operators.Concat(
		"begin-object",
		Ws,
		operators.Terminal("%x7B", []byte{123}),
		Ws,
	)(s)
}

// char = unescaped / escape ( %x22 / %x5C / %x2F / %x62 / %x66 / %x6E / %x72 / %x74 / %x75 4HEXDIG )
func Char(s []byte) operators.Alternatives {
	return operators.Alts(
		"char",
		Unescaped,
		operators.Concat(
			"escape ( %x22 / %x5C / %x2F / %x62 / %x66 / %x6E / %x72 / %x74 / %x75 4HEXDIG )",
			Escape,
			operators.Alts(
				"%x22 / %x5C / %x2F / %x62 / %x66 / %x6E / %x72 / %x74 / %x75 4HEXDIG",
				operators.Terminal("%x22", []byte{34}),
				operators.Terminal("%x5C", []byte{92}),
				operators.Terminal("%x2F", []byte{47}),
				operators.Terminal("%x62", []byte{98}),
				operators.Terminal("%x66", []byte{102}),
				operators.Terminal("%x6E", []byte{110}),
				operators.Terminal("%x72", []byte{114}),
				operators.Terminal("%x74", []byte{116}),
				operators.Concat(
					"%x75 4HEXDIG",
					operators.Terminal("%x75", []byte{117}),
					operators.RepeatN("4HEXDIG", 4, core.HEXDIG()),
				),
			),
		),
	)(s)
}

// decimal-point = %x2E
func DecimalPoint(s []byte) operators.Alternatives {
	return operators.Terminal("decimal-point", []byte{46})(s)
}

// digit1-9 = %x31-39
func Digit19(s []byte) operators.Alternatives {
	return operators.Range("digit1-9", []byte{49}, []byte{57})(s)
}

// e = %x65 / %x45
func E(s []byte) operators.Alternatives {
	return operators.Alts(
		"e",
		operators.Terminal("%x65", []byte{101}),
		operators.Terminal("%x45", []byte{69}),
	)(s)
}

// end-array = ws %x5D ws
func EndArray(s []byte) operators.Alternatives {
	return operators.Concat(
		"end-array",
		Ws,
		operators.Terminal("%x5D", []byte{93}),
		Ws,
	)(s)
}

// end-object = ws %x7D ws
func EndObject(s []byte) operators.Alternatives {
	return operators.Concat(
		"end-object",
		Ws,
		operators.Terminal("%x7D", []byte{125}),
		Ws,
	)(s)
}

// escape = %x5C
func Escape(s []byte) operators.Alternatives {
	return operators.Terminal("escape", []byte{92})(s)
}

// exp = e [ minus / plus ] 1*DIGIT
func Exp(s []byte) operators.Alternatives {
	return operators.Concat(
		"exp",
		E,
		operators.Optional("[ minus / plus ]", operators.Alts(
			"minus / plus",
			Minus,
			Plus,
		)),
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
	)(s)
}

// false = %x66.61.6c.73.65
func False(s []byte) operators.Alternatives {
	return operators.String("false", "false")(s)
}

// frac = decimal-point 1*DIGIT
func Frac(s []byte) operators.Alternatives {
	return operators.Concat(
		"frac",
		DecimalPoint,
		operators.Repeat1Inf("1*DIGIT", core.DIGIT()),
	)(s)
}

// int = zero / ( digit1-9 *DIGIT )
func Int(s []byte) operators.Alternatives {
	return operators.Alts(
		"int",
		Zero,
		operators.Concat(
			"digit1-9 *DIGIT",
			Digit19,
			operators.Repeat0Inf("*DIGIT", core.DIGIT()),
		),
	)(s)
}

// member = string name-separator value
func Member(s []byte) operators.Alternatives {
	return operators.Concat(
		"member",
		String,
		NameSeparator,
		Value,
	)(s)
}

// minus = %x2D
func Minus(s []byte) operators.Alternatives {
	return operators.Terminal("minus", []byte{45})(s)
}

// name-separator = ws %x3A ws
func NameSeparator(s []byte) operators.Alternatives {
	return operators.Concat(
		"name-separator",
		Ws,
		operators.Terminal("%x3A", []byte{58}),
		Ws,
	)(s)
}

// null = %x6e.75.6c.6c
func Null(s []byte) operators.Alternatives {
	return operators.String("null", "null")(s)
}

// number = [ minus ] int [ frac ] [ exp ]
func Number(s []byte) operators.Alternatives {
	return operators.Concat(
		"number",
		operators.Optional("[ minus ]", Minus),
		Int,
		operators.Optional("[ frac ]", Frac),
		operators.Optional("[ exp ]", Exp),
	)(s)
}

// object = begin-object [ member *( value-separator member ) ] end-object
func Object(s []byte) operators.Alternatives {
	return operators.Concat(
		"object",
		BeginObject,
		operators.Optional("[ member *( value-separator member ) ]", operators.Concat(
			"member *( value-separator member )",
			Member,
			operators.Repeat0Inf("*( value-separator member )", operators.Concat(
				"value-separator member",
				ValueSeparator,
				Member,
			)),
		)),
		EndObject,
	)(s)
}

// plus = %x2B
func Plus(s []byte) operators.Alternatives {
	return operators.Terminal("plus", []byte{43})(s)
}

// quotation-mark = %x22
func QuotationMark(s []byte) operators.Alternatives {
	return operators.Terminal("quotation-mark", []byte{34})(s)
}

// string = quotation-mark *char quotation-mark
func String(s []byte) operators.Alternatives {
	return operators.Concat(
		"string",
		QuotationMark,
		operators.Repeat0Inf("*char", Char),
		QuotationMark,
	)(s)
}

// true = %x74.72.75.65
func True(s []byte) operators.Alternatives {
	return operators.String("true", "true")(s)
}

// unescaped = %x20-21 / %x23-5B / %x5D-10FFFF
func Unescaped(s []byte) operators.Alternatives {
	return operators.Alts(
		"unescaped",
		operators.Range("%x20-21", []byte{32}, []byte{33}),
		operators.Range("%x23-5B", []byte{35}, []byte{91}),
		operators.Range("%x5D-10FFFF", []byte{93}, []byte{16, 255, 255}),
	)(s)
}

// value = false / null / true / object / array / number / string
func Value(s []byte) operators.Alternatives {
	return operators.Alts(
		"value",
		False,
		Null,
		True,
		Object,
		Array,
		Number,
		String,
	)(s)
}

// value-separator = ws %x2C ws
func ValueSeparator(s []byte) operators.Alternatives {
	return operators.Concat(
		"value-separator",
		Ws,
		operators.Terminal("%x2C", []byte{44}),
		Ws,
	)(s)
}

// ws = *( %x20 / %x09 / %x0A / %x0D )
func Ws(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("ws", operators.Alts(
		"%x20 / %x09 / %x0A / %x0D",
		operators.Terminal("%x20", []byte{32}),
		operators.Terminal("%x09", []byte{9}),
		operators.Terminal("%x0A", []byte{10}),
		operators.Terminal("%x0D", []byte{13}),
	))(s)
}

// zero = %x30
func Zero(s []byte) operators.Alternatives {
	return operators.Terminal("zero", []byte{48})(s)
}
