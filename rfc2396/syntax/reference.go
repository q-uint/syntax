package syntax

import . "github.com/elimity-com/abnf/operators"

func uRIReference(s []rune) Alternatives {
	return Concat(`URI-reference`,
		Optional(`[ absoluteURI | relativeURI ]`,
			Alts(`absoluteURI | relativeURI`,
				absoluteURI,
				relativeURI,
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

func fragment(s []rune) Alternatives {
	return Repeat0Inf(`fragment`,
		uric,
	)(s)
}
