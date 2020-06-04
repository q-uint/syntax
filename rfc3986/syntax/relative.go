package syntax

import . "github.com/elimity-com/abnf/operators"

func relativeRef(s []rune) Alternatives {
	return Concat(`relative-ref`,
		relativePart,
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

func relativePart(s []rune) Alternatives {
	return Alts(`relative-part`,
		Concat(`"//" authority path-abempty`,
			String(`//`, "//"),
			authority,
			pathAbempty,
		),
		pathAbsolute,
		pathNoscheme,
		pathEmpty,
	)(s)
}
