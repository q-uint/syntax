package syntax

import . "github.com/di-wu/abnf"

func uRIReference(s []rune) *AST {
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

func fragment(s []rune) *AST {
	return Repeat0Inf(`fragment`,
		uric,
	)(s)
}
