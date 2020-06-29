// This file is generated - do not edit.

package syntax

import (
	core "github.com/elimity-com/abnf/core"
	"github.com/elimity-com/abnf/operators"
)

// IP-literal = "[" ( IPv6address / IPvFuture ) "]"
func IPLiteral(s []byte) operators.Alternatives {
	return operators.Concat(
		"IP-literal",
		operators.String("[", "["),
		operators.Alts(
			"IPv6address / IPvFuture",
			IPv6address,
			IPvFuture,
		),
		operators.String("]", "]"),
	)(s)
}

// IPv4address = dec-octet "." dec-octet "." dec-octet "." dec-octet
func IPv4address(s []byte) operators.Alternatives {
	return operators.Concat(
		"IPv4address",
		DecOctet,
		operators.String(".", "."),
		DecOctet,
		operators.String(".", "."),
		DecOctet,
		operators.String(".", "."),
		DecOctet,
	)(s)
}

// IPv6address = 6( h16 ":" ) ls32 / "::" 5( h16 ":" ) ls32 / [ h16 ] "::" 4( h16 ":" ) ls32 / [ *1( h16 ":" ) h16 ] "::" 3( h16 ":" ) ls32 / [ *2( h16 ":" ) h16 ] "::" 2( h16 ":" ) ls32 / [ *3( h16 ":" ) h16 ] "::" h16 ":" ls32 / [ *4( h16 ":" ) h16 ] "::" ls32 / [ *5( h16 ":" ) h16 ] "::" h16 / [ *6( h16 ":" ) h16 ] "::"
func IPv6address(s []byte) operators.Alternatives {
	return operators.Alts(
		"IPv6address",
		operators.Concat(
			"6( h16 \":\" ) ls32",
			operators.Repeat0Inf("6( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.String(":", ":"),
			)),
			Ls32,
		),
		operators.Concat(
			"\"::\" 5( h16 \":\" ) ls32",
			operators.String("::", "::"),
			operators.Repeat0Inf("5( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.String(":", ":"),
			)),
			Ls32,
		),
		operators.Concat(
			"[ h16 ] \"::\" 4( h16 \":\" ) ls32",
			operators.Optional("[ h16 ]", H16),
			operators.String("::", "::"),
			operators.Repeat0Inf("4( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.String(":", ":"),
			)),
			Ls32,
		),
		operators.Concat(
			"[ *1( h16 \":\" ) h16 ] \"::\" 3( h16 \":\" ) ls32",
			operators.Optional("[ *1( h16 \":\" ) h16 ]", operators.Concat(
				"*1( h16 \":\" ) h16",
				operators.Repeat("*1( h16 \":\" )", 0, 1, operators.Concat(
					"h16 \":\"",
					H16,
					operators.String(":", ":"),
				)),
				H16,
			)),
			operators.String("::", "::"),
			operators.Repeat0Inf("3( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.String(":", ":"),
			)),
			Ls32,
		),
		operators.Concat(
			"[ *2( h16 \":\" ) h16 ] \"::\" 2( h16 \":\" ) ls32",
			operators.Optional("[ *2( h16 \":\" ) h16 ]", operators.Concat(
				"*2( h16 \":\" ) h16",
				operators.Repeat("*2( h16 \":\" )", 0, 2, operators.Concat(
					"h16 \":\"",
					H16,
					operators.String(":", ":"),
				)),
				H16,
			)),
			operators.String("::", "::"),
			operators.Repeat0Inf("2( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.String(":", ":"),
			)),
			Ls32,
		),
		operators.Concat(
			"[ *3( h16 \":\" ) h16 ] \"::\" h16 \":\" ls32",
			operators.Optional("[ *3( h16 \":\" ) h16 ]", operators.Concat(
				"*3( h16 \":\" ) h16",
				operators.Repeat("*3( h16 \":\" )", 0, 3, operators.Concat(
					"h16 \":\"",
					H16,
					operators.String(":", ":"),
				)),
				H16,
			)),
			operators.String("::", "::"),
			H16,
			operators.String(":", ":"),
			Ls32,
		),
		operators.Concat(
			"[ *4( h16 \":\" ) h16 ] \"::\" ls32",
			operators.Optional("[ *4( h16 \":\" ) h16 ]", operators.Concat(
				"*4( h16 \":\" ) h16",
				operators.Repeat("*4( h16 \":\" )", 0, 4, operators.Concat(
					"h16 \":\"",
					H16,
					operators.String(":", ":"),
				)),
				H16,
			)),
			operators.String("::", "::"),
			Ls32,
		),
		operators.Concat(
			"[ *5( h16 \":\" ) h16 ] \"::\" h16",
			operators.Optional("[ *5( h16 \":\" ) h16 ]", operators.Concat(
				"*5( h16 \":\" ) h16",
				operators.Repeat("*5( h16 \":\" )", 0, 5, operators.Concat(
					"h16 \":\"",
					H16,
					operators.String(":", ":"),
				)),
				H16,
			)),
			operators.String("::", "::"),
			H16,
		),
		operators.Concat(
			"[ *6( h16 \":\" ) h16 ] \"::\"",
			operators.Optional("[ *6( h16 \":\" ) h16 ]", operators.Concat(
				"*6( h16 \":\" ) h16",
				operators.Repeat("*6( h16 \":\" )", 0, 6, operators.Concat(
					"h16 \":\"",
					H16,
					operators.String(":", ":"),
				)),
				H16,
			)),
			operators.String("::", "::"),
		),
	)(s)
}

// IPvFuture = "v" 1*HEXDIG "." 1*( unreserved / sub-delims / ":" )
func IPvFuture(s []byte) operators.Alternatives {
	return operators.Concat(
		"IPvFuture",
		operators.String("v", "v"),
		operators.Repeat1Inf("1*HEXDIG", core.HEXDIG()),
		operators.String(".", "."),
		operators.Repeat1Inf("1*( unreserved / sub-delims / \":\" )", operators.Alts(
			"unreserved / sub-delims / \":\"",
			Unreserved,
			SubDelims,
			operators.String(":", ":"),
		)),
	)(s)
}

// URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ]
func URI(s []byte) operators.Alternatives {
	return operators.Concat(
		"URI",
		Scheme,
		operators.String(":", ":"),
		HierPart,
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.String("?", "?"),
			Query,
		)),
		operators.Optional("[ \"#\" fragment ]", operators.Concat(
			"\"#\" fragment",
			operators.String("#", "#"),
			Fragment,
		)),
	)(s)
}

// URI-reference = URI / relative-ref
func URIReference(s []byte) operators.Alternatives {
	return operators.Alts(
		"URI-reference",
		URI,
		RelativeRef,
	)(s)
}

// absolute-URI = scheme ":" hier-part [ "?" query ]
func AbsoluteURI(s []byte) operators.Alternatives {
	return operators.Concat(
		"absolute-URI",
		Scheme,
		operators.String(":", ":"),
		HierPart,
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.String("?", "?"),
			Query,
		)),
	)(s)
}

// authority = [ userinfo "@" ] host [ ":" port ]
func Authority(s []byte) operators.Alternatives {
	return operators.Concat(
		"authority",
		operators.Optional("[ userinfo \"@\" ]", operators.Concat(
			"userinfo \"@\"",
			Userinfo,
			operators.String("@", "@"),
		)),
		Host,
		operators.Optional("[ \":\" port ]", operators.Concat(
			"\":\" port",
			operators.String(":", ":"),
			Port,
		)),
	)(s)
}

// dec-octet = DIGIT ; 0-9 / %x31-39 DIGIT ; 10-99 / "1" 2DIGIT ; 100-199 / "2" %x30-34 DIGIT ; 200-249 / "25" %x30-35
func DecOctet(s []byte) operators.Alternatives {
	return operators.Alts(
		"dec-octet",
		core.DIGIT(),
		operators.Concat(
			"%x31-39 DIGIT",
			operators.Range("%x31-39", []byte{49}, []byte{57}),
			core.DIGIT(),
		),
		operators.Concat(
			"\"1\" 2DIGIT",
			operators.String("1", "1"),
			operators.Repeat0Inf("2DIGIT", core.DIGIT()),
		),
		operators.Concat(
			"\"2\" %x30-34 DIGIT",
			operators.String("2", "2"),
			operators.Range("%x30-34", []byte{48}, []byte{52}),
			core.DIGIT(),
		),
		operators.Concat(
			"\"25\" %x30-35",
			operators.String("25", "25"),
			operators.Range("%x30-35", []byte{48}, []byte{53}),
		),
	)(s)
}

// fragment = *( pchar / "/" / "?" )
func Fragment(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("fragment", operators.Alts(
		"pchar / \"/\" / \"?\"",
		Pchar,
		operators.String("/", "/"),
		operators.String("?", "?"),
	))(s)
}

// gen-delims = ":" / "/" / "?" / "#" / "[" / "]" / "@"
func GenDelims(s []byte) operators.Alternatives {
	return operators.Alts(
		"gen-delims",
		operators.String(":", ":"),
		operators.String("/", "/"),
		operators.String("?", "?"),
		operators.String("#", "#"),
		operators.String("[", "["),
		operators.String("]", "]"),
		operators.String("@", "@"),
	)(s)
}

// h16 = 1*4HEXDIG
func H16(s []byte) operators.Alternatives {
	return operators.Repeat("h16", 1, 4, core.HEXDIG())(s)
}

// hier-part = "//" authority path-abempty / path-absolute / path-rootless / path-empty
func HierPart(s []byte) operators.Alternatives {
	return operators.Alts(
		"hier-part",
		operators.Concat(
			"\"//\" authority path-abempty",
			operators.String("//", "//"),
			Authority,
			PathAbempty,
		),
		PathAbsolute,
		PathRootless,
		PathEmpty,
	)(s)
}

// host = IP-literal / IPv4address / reg-name
func Host(s []byte) operators.Alternatives {
	return operators.Alts(
		"host",
		IPLiteral,
		IPv4address,
		RegName,
	)(s)
}

// ls32 = ( h16 ":" h16 ) / IPv4address
func Ls32(s []byte) operators.Alternatives {
	return operators.Alts(
		"ls32",
		operators.Concat(
			"h16 \":\" h16",
			H16,
			operators.String(":", ":"),
			H16,
		),
		IPv4address,
	)(s)
}

// path = path-abempty ; begins with "/" or is empty / path-absolute ; begins with "/" but not "//" / path-noscheme ; begins with a non-colon segment / path-rootless ; begins with a segment / path-empty
func Path(s []byte) operators.Alternatives {
	return operators.Alts(
		"path",
		PathAbempty,
		PathAbsolute,
		PathNoscheme,
		PathRootless,
		PathEmpty,
	)(s)
}

// path-abempty = *( "/" segment )
func PathAbempty(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("path-abempty", operators.Concat(
		"\"/\" segment",
		operators.String("/", "/"),
		Segment,
	))(s)
}

// path-absolute = "/" [ segment-nz *( "/" segment ) ]
func PathAbsolute(s []byte) operators.Alternatives {
	return operators.Concat(
		"path-absolute",
		operators.String("/", "/"),
		operators.Optional("[ segment-nz *( \"/\" segment ) ]", operators.Concat(
			"segment-nz *( \"/\" segment )",
			SegmentNz,
			operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
				"\"/\" segment",
				operators.String("/", "/"),
				Segment,
			)),
		)),
	)(s)
}

// path-empty = 0pchar
func PathEmpty(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("path-empty", Pchar)(s)
}

// path-noscheme = segment-nz-nc *( "/" segment )
func PathNoscheme(s []byte) operators.Alternatives {
	return operators.Concat(
		"path-noscheme",
		SegmentNzNc,
		operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
			"\"/\" segment",
			operators.String("/", "/"),
			Segment,
		)),
	)(s)
}

// path-rootless = segment-nz *( "/" segment )
func PathRootless(s []byte) operators.Alternatives {
	return operators.Concat(
		"path-rootless",
		SegmentNz,
		operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
			"\"/\" segment",
			operators.String("/", "/"),
			Segment,
		)),
	)(s)
}

// pchar = unreserved / pct-encoded / sub-delims / ":" / "@"
func Pchar(s []byte) operators.Alternatives {
	return operators.Alts(
		"pchar",
		Unreserved,
		PctEncoded,
		SubDelims,
		operators.String(":", ":"),
		operators.String("@", "@"),
	)(s)
}

// pct-encoded = "%" HEXDIG HEXDIG
func PctEncoded(s []byte) operators.Alternatives {
	return operators.Concat(
		"pct-encoded",
		operators.String("%", "%"),
		core.HEXDIG(),
		core.HEXDIG(),
	)(s)
}

// port = *DIGIT
func Port(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("port", core.DIGIT())(s)
}

// query = *( pchar / "/" / "?" )
func Query(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("query", operators.Alts(
		"pchar / \"/\" / \"?\"",
		Pchar,
		operators.String("/", "/"),
		operators.String("?", "?"),
	))(s)
}

// reg-name = *( unreserved / pct-encoded / sub-delims )
func RegName(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("reg-name", operators.Alts(
		"unreserved / pct-encoded / sub-delims",
		Unreserved,
		PctEncoded,
		SubDelims,
	))(s)
}

// relative-part = "//" authority path-abempty / path-absolute / path-noscheme / path-empty
func RelativePart(s []byte) operators.Alternatives {
	return operators.Alts(
		"relative-part",
		operators.Concat(
			"\"//\" authority path-abempty",
			operators.String("//", "//"),
			Authority,
			PathAbempty,
		),
		PathAbsolute,
		PathNoscheme,
		PathEmpty,
	)(s)
}

// relative-ref = relative-part [ "?" query ] [ "#" fragment ]
func RelativeRef(s []byte) operators.Alternatives {
	return operators.Concat(
		"relative-ref",
		RelativePart,
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.String("?", "?"),
			Query,
		)),
		operators.Optional("[ \"#\" fragment ]", operators.Concat(
			"\"#\" fragment",
			operators.String("#", "#"),
			Fragment,
		)),
	)(s)
}

// reserved = gen-delims / sub-delims
func Reserved(s []byte) operators.Alternatives {
	return operators.Alts(
		"reserved",
		GenDelims,
		SubDelims,
	)(s)
}

// scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
func Scheme(s []byte) operators.Alternatives {
	return operators.Concat(
		"scheme",
		core.ALPHA(),
		operators.Repeat0Inf("*( ALPHA / DIGIT / \"+\" / \"-\" / \".\" )", operators.Alts(
			"ALPHA / DIGIT / \"+\" / \"-\" / \".\"",
			core.ALPHA(),
			core.DIGIT(),
			operators.String("+", "+"),
			operators.String("-", "-"),
			operators.String(".", "."),
		)),
	)(s)
}

// segment = *pchar
func Segment(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("segment", Pchar)(s)
}

// segment-nz = 1*pchar
func SegmentNz(s []byte) operators.Alternatives {
	return operators.Repeat1Inf("segment-nz", Pchar)(s)
}

// segment-nz-nc = 1*( unreserved / pct-encoded / sub-delims / "@" )
func SegmentNzNc(s []byte) operators.Alternatives {
	return operators.Repeat1Inf("segment-nz-nc", operators.Alts(
		"unreserved / pct-encoded / sub-delims / \"@\"",
		Unreserved,
		PctEncoded,
		SubDelims,
		operators.String("@", "@"),
	))(s)
}

// sub-delims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
func SubDelims(s []byte) operators.Alternatives {
	return operators.Alts(
		"sub-delims",
		operators.String("!", "!"),
		operators.String("$", "$"),
		operators.String("&", "&"),
		operators.String("'", "'"),
		operators.String("(", "("),
		operators.String(")", ")"),
		operators.String("*", "*"),
		operators.String("+", "+"),
		operators.String(",", ","),
		operators.String(";", ";"),
		operators.String("=", "="),
	)(s)
}

// unreserved = ALPHA / DIGIT / "-" / "." / "_" / "~"
func Unreserved(s []byte) operators.Alternatives {
	return operators.Alts(
		"unreserved",
		core.ALPHA(),
		core.DIGIT(),
		operators.String("-", "-"),
		operators.String(".", "."),
		operators.String("_", "_"),
		operators.String("~", "~"),
	)(s)
}

// userinfo = *( unreserved / pct-encoded / sub-delims / ":" )
func Userinfo(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("userinfo", operators.Alts(
		"unreserved / pct-encoded / sub-delims / \":\"",
		Unreserved,
		PctEncoded,
		SubDelims,
		operators.String(":", ":"),
	))(s)
}
