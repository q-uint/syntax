// This file is generated - do not edit.

package syntax

import "github.com/elimity-com/abnf/operators"

// IPv4address = 1*digit "." 1*digit "." 1*digit "." 1*digit
func IPv4address(s []byte) operators.Alternatives {
	return operators.Concat(
		"IPv4address",
		operators.Repeat1Inf("1*digit", Digit),
		operators.String(".", "."),
		operators.Repeat1Inf("1*digit", Digit),
		operators.String(".", "."),
		operators.Repeat1Inf("1*digit", Digit),
		operators.String(".", "."),
		operators.Repeat1Inf("1*digit", Digit),
	)(s)
}

// URI-reference = [ absoluteURI / relativeURI ] [ "#" fragment ]
func URIReference(s []byte) operators.Alternatives {
	return operators.Concat(
		"URI-reference",
		operators.Optional("[ absoluteURI / relativeURI ]", operators.Alts(
			"absoluteURI / relativeURI",
			AbsoluteURI,
			RelativeURI,
		)),
		operators.Optional("[ \"#\" fragment ]", operators.Concat(
			"\"#\" fragment",
			operators.String("#", "#"),
			Fragment,
		)),
	)(s)
}

// abs-path = "/" path-segments
func AbsPath(s []byte) operators.Alternatives {
	return operators.Concat(
		"abs-path",
		operators.String("/", "/"),
		PathSegments,
	)(s)
}

// absoluteURI = scheme ":" ( hier-part / opaque-part )
func AbsoluteURI(s []byte) operators.Alternatives {
	return operators.Concat(
		"absoluteURI",
		Scheme,
		operators.String(":", ":"),
		operators.Alts(
			"hier-part / opaque-part",
			HierPart,
			OpaquePart,
		),
	)(s)
}

// alpha = lowalpha / upalpha
func Alpha(s []byte) operators.Alternatives {
	return operators.Alts(
		"alpha",
		Lowalpha,
		Upalpha,
	)(s)
}

// alphanum = alpha / digit
func Alphanum(s []byte) operators.Alternatives {
	return operators.Alts(
		"alphanum",
		Alpha,
		Digit,
	)(s)
}

// authority = server / reg-name
func Authority(s []byte) operators.Alternatives {
	return operators.Alts(
		"authority",
		Server,
		RegName,
	)(s)
}

// digit = "0" / "1" / "2" / "3" / "4" / "5" / "6" / "7" / "8" / "9"
func Digit(s []byte) operators.Alternatives {
	return operators.Alts(
		"digit",
		operators.String("0", "0"),
		operators.String("1", "1"),
		operators.String("2", "2"),
		operators.String("3", "3"),
		operators.String("4", "4"),
		operators.String("5", "5"),
		operators.String("6", "6"),
		operators.String("7", "7"),
		operators.String("8", "8"),
		operators.String("9", "9"),
	)(s)
}

// domainlabel = alphanum / alphanum *( alphanum / "-" ) alphanum
func Domainlabel(s []byte) operators.Alternatives {
	return operators.Alts(
		"domainlabel",
		Alphanum,
		operators.Concat(
			"alphanum *( alphanum / \"-\" ) alphanum",
			Alphanum,
			operators.Repeat0Inf("*( alphanum / \"-\" )", operators.Alts(
				"alphanum / \"-\"",
				Alphanum,
				operators.String("-", "-"),
			)),
			Alphanum,
		),
	)(s)
}

// escaped = "%" hex hex
func Escaped(s []byte) operators.Alternatives {
	return operators.Concat(
		"escaped",
		operators.String("%", "%"),
		Hex,
		Hex,
	)(s)
}

// fragment = *uric
func Fragment(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("fragment", Uric)(s)
}

// hex = digit / "A" / "B" / "C" / "D" / "E" / "F" / "a" / "b" / "c" / "d" / "e" / "f"
func Hex(s []byte) operators.Alternatives {
	return operators.Alts(
		"hex",
		Digit,
		operators.String("A", "A"),
		operators.String("B", "B"),
		operators.String("C", "C"),
		operators.String("D", "D"),
		operators.String("E", "E"),
		operators.String("F", "F"),
		operators.String("a", "a"),
		operators.String("b", "b"),
		operators.String("c", "c"),
		operators.String("d", "d"),
		operators.String("e", "e"),
		operators.String("f", "f"),
	)(s)
}

// hier-part = ( net-path / abs-path ) [ "?" query ]
func HierPart(s []byte) operators.Alternatives {
	return operators.Concat(
		"hier-part",
		operators.Alts(
			"net-path / abs-path",
			NetPath,
			AbsPath,
		),
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.String("?", "?"),
			Query,
		)),
	)(s)
}

// host = hostname / IPv4address
func Host(s []byte) operators.Alternatives {
	return operators.Alts(
		"host",
		Hostname,
		IPv4address,
	)(s)
}

// hostname = *( domainlabel "." ) toplabel [ "." ]
func Hostname(s []byte) operators.Alternatives {
	return operators.Concat(
		"hostname",
		operators.Repeat0Inf("*( domainlabel \".\" )", operators.Concat(
			"domainlabel \".\"",
			Domainlabel,
			operators.String(".", "."),
		)),
		Toplabel,
		operators.Optional("[ \".\" ]", operators.String(".", ".")),
	)(s)
}

// hostport = host [ ":" port ]
func Hostport(s []byte) operators.Alternatives {
	return operators.Concat(
		"hostport",
		Host,
		operators.Optional("[ \":\" port ]", operators.Concat(
			"\":\" port",
			operators.String(":", ":"),
			Port,
		)),
	)(s)
}

// lowalpha = "a" / "b" / "c" / "d" / "e" / "f" / "g" / "h" / "i" / "j" / "k" / "l" / "m" / "n" / "o" / "p" / "q" / "r" / "s" / "t" / "u" / "v" / "w" / "x" / "y" / "z"
func Lowalpha(s []byte) operators.Alternatives {
	return operators.Alts(
		"lowalpha",
		operators.String("a", "a"),
		operators.String("b", "b"),
		operators.String("c", "c"),
		operators.String("d", "d"),
		operators.String("e", "e"),
		operators.String("f", "f"),
		operators.String("g", "g"),
		operators.String("h", "h"),
		operators.String("i", "i"),
		operators.String("j", "j"),
		operators.String("k", "k"),
		operators.String("l", "l"),
		operators.String("m", "m"),
		operators.String("n", "n"),
		operators.String("o", "o"),
		operators.String("p", "p"),
		operators.String("q", "q"),
		operators.String("r", "r"),
		operators.String("s", "s"),
		operators.String("t", "t"),
		operators.String("u", "u"),
		operators.String("v", "v"),
		operators.String("w", "w"),
		operators.String("x", "x"),
		operators.String("y", "y"),
		operators.String("z", "z"),
	)(s)
}

// mark = "-" / "-" / "." / "!" / "~" / "*" / "'" / "(" / ")"
func Mark(s []byte) operators.Alternatives {
	return operators.Alts(
		"mark",
		operators.String("-", "-"),
		operators.String("-", "-"),
		operators.String(".", "."),
		operators.String("!", "!"),
		operators.String("~", "~"),
		operators.String("*", "*"),
		operators.String("'", "'"),
		operators.String("(", "("),
		operators.String(")", ")"),
	)(s)
}

// net-path = "//" authority [ abs-path ]
func NetPath(s []byte) operators.Alternatives {
	return operators.Concat(
		"net-path",
		operators.String("//", "//"),
		Authority,
		operators.Optional("[ abs-path ]", AbsPath),
	)(s)
}

// opaque-part = uric-no-slash *uric
func OpaquePart(s []byte) operators.Alternatives {
	return operators.Concat(
		"opaque-part",
		UricNoSlash,
		operators.Repeat0Inf("*uric", Uric),
	)(s)
}

// param = *pchar
func Param(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("param", Pchar)(s)
}

// path = [ abs-path / opaque-part ]
func Path(s []byte) operators.Alternatives {
	return operators.Optional("path", operators.Alts(
		"abs-path / opaque-part",
		AbsPath,
		OpaquePart,
	))(s)
}

// path-segments = segment *( "/" segment )
func PathSegments(s []byte) operators.Alternatives {
	return operators.Concat(
		"path-segments",
		Segment,
		operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
			"\"/\" segment",
			operators.String("/", "/"),
			Segment,
		)),
	)(s)
}

// pchar = unreserved / escaped / ":" / "@" / "&" / "=" / "+" / "$" / ","
func Pchar(s []byte) operators.Alternatives {
	return operators.Alts(
		"pchar",
		Unreserved,
		Escaped,
		operators.String(":", ":"),
		operators.String("@", "@"),
		operators.String("&", "&"),
		operators.String("=", "="),
		operators.String("+", "+"),
		operators.String("$", "$"),
		operators.String(",", ","),
	)(s)
}

// port = *digit
func Port(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("port", Digit)(s)
}

// query = *uric
func Query(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("query", Uric)(s)
}

// reg-name = 1*( unreserved / escaped / "$" / "," / ";" / ":" / "@" / "&" / "=" / "+" )
func RegName(s []byte) operators.Alternatives {
	return operators.Repeat1Inf("reg-name", operators.Alts(
		"unreserved / escaped / \"$\" / \",\" / \";\" / \":\" / \"@\" / \"&\" / \"=\" / \"+\"",
		Unreserved,
		Escaped,
		operators.String("$", "$"),
		operators.String(",", ","),
		operators.String(";", ";"),
		operators.String(":", ":"),
		operators.String("@", "@"),
		operators.String("&", "&"),
		operators.String("=", "="),
		operators.String("+", "+"),
	))(s)
}

// rel-path = rel-segment [ abs-path ]
func RelPath(s []byte) operators.Alternatives {
	return operators.Concat(
		"rel-path",
		RelSegment,
		operators.Optional("[ abs-path ]", AbsPath),
	)(s)
}

// rel-segment = 1*( unreserved / escaped / ";" / "@" / "&" / "=" / "+" / "$" / "," )
func RelSegment(s []byte) operators.Alternatives {
	return operators.Repeat1Inf("rel-segment", operators.Alts(
		"unreserved / escaped / \";\" / \"@\" / \"&\" / \"=\" / \"+\" / \"$\" / \",\"",
		Unreserved,
		Escaped,
		operators.String(";", ";"),
		operators.String("@", "@"),
		operators.String("&", "&"),
		operators.String("=", "="),
		operators.String("+", "+"),
		operators.String("$", "$"),
		operators.String(",", ","),
	))(s)
}

// relativeURI = ( net-path / abs-path / rel-path ) [ "?" query ]
func RelativeURI(s []byte) operators.Alternatives {
	return operators.Concat(
		"relativeURI",
		operators.Alts(
			"net-path / abs-path / rel-path",
			NetPath,
			AbsPath,
			RelPath,
		),
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.String("?", "?"),
			Query,
		)),
	)(s)
}

// reserved = ";" / "/" / "?" / ":" / "@" / "&" / "=" / "+" / "$" / ","
func Reserved(s []byte) operators.Alternatives {
	return operators.Alts(
		"reserved",
		operators.String(";", ";"),
		operators.String("/", "/"),
		operators.String("?", "?"),
		operators.String(":", ":"),
		operators.String("@", "@"),
		operators.String("&", "&"),
		operators.String("=", "="),
		operators.String("+", "+"),
		operators.String("$", "$"),
		operators.String(",", ","),
	)(s)
}

// scheme = alpha *( alpha / digit / "+" / "-" / "." )
func Scheme(s []byte) operators.Alternatives {
	return operators.Concat(
		"scheme",
		Alpha,
		operators.Repeat0Inf("*( alpha / digit / \"+\" / \"-\" / \".\" )", operators.Alts(
			"alpha / digit / \"+\" / \"-\" / \".\"",
			Alpha,
			Digit,
			operators.String("+", "+"),
			operators.String("-", "-"),
			operators.String(".", "."),
		)),
	)(s)
}

// segment = *pchar *( ";" param )
func Segment(s []byte) operators.Alternatives {
	return operators.Concat(
		"segment",
		operators.Repeat0Inf("*pchar", Pchar),
		operators.Repeat0Inf("*( \";\" param )", operators.Concat(
			"\";\" param",
			operators.String(";", ";"),
			Param,
		)),
	)(s)
}

// server = [ [ userinfo "@" ] hostport ]
func Server(s []byte) operators.Alternatives {
	return operators.Optional("server", operators.Concat(
		"[ userinfo \"@\" ] hostport",
		operators.Optional("[ userinfo \"@\" ]", operators.Concat(
			"userinfo \"@\"",
			Userinfo,
			operators.String("@", "@"),
		)),
		Hostport,
	))(s)
}

// toplabel = alpha / alpha *( alphanum / "-" ) alphanum
func Toplabel(s []byte) operators.Alternatives {
	return operators.Alts(
		"toplabel",
		Alpha,
		operators.Concat(
			"alpha *( alphanum / \"-\" ) alphanum",
			Alpha,
			operators.Repeat0Inf("*( alphanum / \"-\" )", operators.Alts(
				"alphanum / \"-\"",
				Alphanum,
				operators.String("-", "-"),
			)),
			Alphanum,
		),
	)(s)
}

// unreserved = alphanum / mark
func Unreserved(s []byte) operators.Alternatives {
	return operators.Alts(
		"unreserved",
		Alphanum,
		Mark,
	)(s)
}

// upalpha = "A" / "B" / "C" / "D" / "E" / "F" / "G" / "H" / "I" / "J" / "K" / "L" / "M" / "N" / "O" / "P" / "Q" / "R" / "S" / "T" / "U" / "V" / "W" / "X" / "Y" / "Z"
func Upalpha(s []byte) operators.Alternatives {
	return operators.Alts(
		"upalpha",
		operators.String("A", "A"),
		operators.String("B", "B"),
		operators.String("C", "C"),
		operators.String("D", "D"),
		operators.String("E", "E"),
		operators.String("F", "F"),
		operators.String("G", "G"),
		operators.String("H", "H"),
		operators.String("I", "I"),
		operators.String("J", "J"),
		operators.String("K", "K"),
		operators.String("L", "L"),
		operators.String("M", "M"),
		operators.String("N", "N"),
		operators.String("O", "O"),
		operators.String("P", "P"),
		operators.String("Q", "Q"),
		operators.String("R", "R"),
		operators.String("S", "S"),
		operators.String("T", "T"),
		operators.String("U", "U"),
		operators.String("V", "V"),
		operators.String("W", "W"),
		operators.String("X", "X"),
		operators.String("Y", "Y"),
		operators.String("Z", "Z"),
	)(s)
}

// uric = reserved / unreserved / escaped
func Uric(s []byte) operators.Alternatives {
	return operators.Alts(
		"uric",
		Reserved,
		Unreserved,
		Escaped,
	)(s)
}

// uric-no-slash = unreserved / escaped / ";" / "?" / ":" / "@" / "&" / "=" / "+" / "$" / ","
func UricNoSlash(s []byte) operators.Alternatives {
	return operators.Alts(
		"uric-no-slash",
		Unreserved,
		Escaped,
		operators.String(";", ";"),
		operators.String("?", "?"),
		operators.String(":", ":"),
		operators.String("@", "@"),
		operators.String("&", "&"),
		operators.String("=", "="),
		operators.String("+", "+"),
		operators.String("$", "$"),
		operators.String(",", ","),
	)(s)
}

// userinfo = *( unreserved / escaped / ";" / ":" / "&" / "=" / "+" / "$" / "," )
func Userinfo(s []byte) operators.Alternatives {
	return operators.Repeat0Inf("userinfo", operators.Alts(
		"unreserved / escaped / \";\" / \":\" / \"&\" / \"=\" / \"+\" / \"$\" / \",\"",
		Unreserved,
		Escaped,
		operators.String(";", ";"),
		operators.String(":", ":"),
		operators.String("&", "&"),
		operators.String("=", "="),
		operators.String("+", "+"),
		operators.String("$", "$"),
		operators.String(",", ","),
	))(s)
}
