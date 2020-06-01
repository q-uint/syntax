package syntax

import . "github.com/di-wu/abnf"

var (
	// URI Characters and Escape Sequences
	// https://tools.ietf.org/html/rfc2396#section-2
	uric = Alts(`uric`, reserved, unreserved, escaped)

	// Reserved Characters
	// https://tools.ietf.org/html/rfc2396#section-2.2
	reserved = Alts(`reserved`,
		Rune(`;`, ';'),
		Rune(`/`, '/'),
		Rune(`?`, '?'),
		Rune(`:`, ':'),
		Rune(`@`, '@'),
		Rune(`&`, '&'),
		Rune(`=`, '='),
		Rune(`+`, '+'),
		Rune(`$`, '$'),
		Rune(`,`, ','),
	)

	// Unreserved Characters
	// https://tools.ietf.org/html/rfc2396#section-2.3
	unreserved = Alts(`unreserved`, alphanum, mark)
	mark       = Alts(`mark`,
		Rune(`-`, '-'),
		Rune(`_`, '_'),
		Rune(`.`, '.'),
		Rune(`!`, '!'),
		Rune(`~`, '~'),
		Rune(`*`, '*'),
		Rune(`'`, '\''),
		Rune(`(`, '('),
		Rune(`)`, ')'),
	)

	// Escaped Encoding
	// https://tools.ietf.org/html/rfc2396#section-2.4
	escaped = Concat(`escaped`, Rune(`%`, '%'), hex, hex)
	hex     = Alts(`hex`,
		digit,
		Range(`A-F`, 'A', 'F'),
		Range(`a-f`, 'a', 'f'),
	)

	// Excluded US-ASCII Characters
	// https://tools.ietf.org/html/rfc2396#section-2.4.3
	control = Alts(`control`,
		Range(`0x00-1F`, 0, 31),
		Rune(`0x7F`, 127),
	)
	space  = Rune(` `, ' ')
	delims = Alts(`delims`,
		Rune(`<`, '<'),
		Rune(`>`, '>'),
		Rune(`#`, '#'),
		Rune(`%`, '%'),
		Rune(`"`, '"'),
	)
	unwise = Alts(`unwise`,
		Rune(`{`, '{'),
		Rune(`}`, '}'),
		Rune(`|`, '|'),
		Rune(`\`, '\\'),
		Rune(`^`, '^'),
		Rune(`[`, '['),
		Rune(`]`, ']'),
		Rune("`", '`'),
	)
)
