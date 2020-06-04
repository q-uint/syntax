package syntax

import . "github.com/elimity-com/abnf/operators"

var (
	alpha  = Alts(`ALPHA`, Range(`%x41-5A`, '\x41', '\x5A'), Range(`%x61-7A`, '\x61', '\x7A'))
	digit  = Range(`DIGIT`, '\x30', '\x39')
	hexdig = Alts(`HEXDIG`, digit,
		Rune(`A`, 'A'), Rune(`B`, 'B'), Rune(`C`, 'C'),
		Rune(`D`, 'D'), Rune(`E`, 'E'), Rune(`F`, 'F'),
		// not in RFC?
		Rune(`a`, 'a'), Rune(`b`, 'b'), Rune(`c`, 'c'),
		Rune(`d`, 'd'), Rune(`e`, 'e'), Rune(`f`, 'f'),
	)
)
