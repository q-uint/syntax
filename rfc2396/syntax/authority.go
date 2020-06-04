package syntax

import . "github.com/elimity-com/abnf/operators"

func authority(s []rune) Alternatives {
	return Alts(`authority`,
		server,
		regName,
	)(s)
}

func regName(s []rune) Alternatives {
	return Repeat1Inf(`reg_name`,
		Alts(`unreserved | escaped | "$" | "," | ";" | ":" | "@" | "&" | "=" | "+"`,
			unreserved,
			escaped,
			Rune(`$`, '$'),
			Rune(`,`, ','),
			Rune(`;`, ';'),
			Rune(`:`, ':'),
			Rune(`@`, '@'),
			Rune(`&`, '&'),
			Rune(`=`, '='),
			Rune(`+`, '+'),
		),
	)(s)
}

func server(s []rune) Alternatives {
	return Optional(`server`,
		Concat(`[ userinfo "@" ] hostport`,
			Optional(`[ userinfo "@" ]`,
				Concat(`userinfo "@"`,
					userinfo,
					Rune(`@`, '@'),
				),
			),
			hostport,
		),
	)(s)
}

func userinfo(s []rune) Alternatives {
	return Repeat1Inf(`userinfo`,
		Alts(`unreserved | escaped | ";" | ":" | "&" | "=" | "+" | "$" | ","`,
			unreserved,
			escaped,
			Rune(`;`, ';'),
			Rune(`:`, ':'),
			Rune(`&`, '&'),
			Rune(`=`, '='),
			Rune(`+`, '+'),
			Rune(`$`, '$'),
			Rune(`,`, ','),
		),
	)(s)
}

func hostport(s []rune) Alternatives {
	return Concat(`hostport`,
		host,
		Optional(`[ ":" port ]`,
			Concat(`":" port`,
				Rune(`:`, ':'),
				port,
			),
		),
	)(s)
}

func host(s []rune) Alternatives {
	return Alts(`host`,
		hostname,
		iPv4address,
	)(s)
}

func hostname(s []rune) Alternatives {
	return Concat(`hostname`,
		Repeat0Inf(`*( domainlabel "." )`,
			Concat(`domainlabel "."`,
				domainlabel,
				Rune(`.`, '.'),
			),
		),
		toplabel,
		Optional(`[ "." ]`, Rune(`.`, '.')),
	)(s)
}

func domainlabel(s []rune) Alternatives {
	return Alts(`domainlabel`,
		alphanum,
		Concat(`alphanum *( alphanum | "-" ) alphanum`,
			alphanum,
			Repeat0Inf(`*( alphanum | "-" )`,
				Alts(`alphanum | "-"`,
					alphanum,
					Rune(`-`, '-'),
				),
			),
			alphanum,
		),
	)(s)
}

func toplabel(s []rune) Alternatives {
	return Alts(`toplabel`,
		alpha,
		Concat(`alpha *( alphanum | "-" ) alphanum`,
			alpha,
			Repeat0Inf(`*( alphanum | "-" )`,
				Alts(`alphanum | "-"`,
					alphanum,
					Rune(`-`, '-'),
				),
			),
			alphanum,
		),
	)(s)
}

func iPv4address(s []rune) Alternatives {
	return Concat(`IPv4address`,
		Repeat1Inf(`1*digit`, digit),
		Rune(`.`, '.'),
		Repeat1Inf(`1*digit`, digit),
		Rune(`.`, '.'),
		Repeat1Inf(`1*digit`, digit),
		Rune(`.`, '.'),
		Repeat1Inf(`1*digit`, digit),
	)(s)
}

func port(s []rune) Alternatives {
	return Repeat0Inf(`*digit`, digit)(s)
}
