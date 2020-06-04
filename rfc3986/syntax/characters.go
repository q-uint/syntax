package syntax

import . "github.com/elimity-com/abnf/operators"

var (
	// https://tools.ietf.org/html/rfc3986#section-2.1
	pctEncoded = Concat(`escaped`,
		Rune(`%`, '%'),
		hexdig,
		hexdig,
	)

	// https://tools.ietf.org/html/rfc3986#section-2.2
	reserved = Alts(`reserved`,
		genDelims,
		subDelims,
	)
	genDelims = Alts(`gen-delims`,
		Rune(`:`, ':'),
		Rune(`/`, '/'),
		Rune(`?`, '?'),
		Rune(`#`, '#'),
		Rune(`[`, '['),
		Rune(`]`, ']'),
		Rune(`@`, '@'),
	)
	subDelims = Alts(`sub-delims`,
		Rune(`!`, '!'),
		Rune(`$`, '$'),
		Rune(`&`, '&'),
		Rune(`'`, '\''),
		Rune(`(`, '('),
		Rune(`)`, ')'),
		Rune(`*`, '*'),
		Rune(`+`, '+'),
		Rune(`,`, ','),
		Rune(`;`, ';'),
		Rune(`=`, '='),
	)
	unreserved = Alts(`unreserved`,
		alpha,
		digit,
		Rune(`-`, '-'),
		Rune(`.`, '.'),
		Rune(`_`, '_'),
		Rune(`~`, '~'),
	)
)
