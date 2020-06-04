package syntax

import . "github.com/elimity-com/abnf/operators"

func absoluteURI(s []rune) Alternatives {
	return Concat(`absolute-URI`,
		scheme,
		Rune(`:`, ':'),
		hierPart,
		Optional(`[ "?" query ]`,
			Concat(`"?" query`,
				Rune(`?`, '?'),
				query,
			),
		),
	)(s)
}
