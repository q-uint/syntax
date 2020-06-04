package syntax

import . "github.com/elimity-com/abnf/operators"

func uRI(s []rune) Alternatives {
	return Concat(`URI`,
		scheme,
		Rune(`:`, ':'),
		hierPart,
		Optional(`[ "?" query ]`,
			Concat(`"?" query`,
				Rune(`?`, '?'),
				query,
			),
		),
		Optional(`[ "#" fragment ]`,
			Concat(`"#" fragment`,
				Rune(`#`, '#'),
				fragment,
			),
		),
	)(s)
}

func hierPart(s []rune) Alternatives {
	return Alts(`hier-part`,
		Concat(`"//" authority path-abempty`,
			String(`//`, "//"),
			authority,
			pathAbempty,
		),
		pathAbsolute,
		pathRootless,
		pathEmpty,
	)(s)
}

func uRIReference(s []rune) Alternatives {
	return Alts(`URI-reference`,
		uRI,
		relativeRef,
	)(s)
}
