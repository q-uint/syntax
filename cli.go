package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elimity-com/abnf"
)

func main() {
	rfcNumber := flag.Int("number", 0, "the number of the rfc spec. e.g. 2396")
	flag.Parse()

	raw, err := ioutil.ReadFile(fmt.Sprintf("./abnf/rfc%d.abnf", *rfcNumber))
	if err != nil {
		log.Fatal(err)
	}

	g := abnf.CodeGenerator{
		PackageName: "syntax",
		RawABNF:     raw,
	}

	externalABNF, ok := dependencies[*rfcNumber]
	if ok {
		g.ExternalABNF = externalABNF
	}

	f := g.GenerateABNFAsAlternatives()
	if err := ioutil.WriteFile(
		fmt.Sprintf("./rfc%d/syntax.go", *rfcNumber),
		[]byte(fmt.Sprintf("%#v", f)), 0644,
	); err != nil {
		log.Fatal(err)
	}
}

var (
	corePkg = abnf.ExternalABNF{
		IsOperator:  true,
		PackageName: "github.com/elimity-com/abnf/core",
	}

	dependencies = map[int]map[string]abnf.ExternalABNF{
		3986: {
			"ALPHA":  corePkg,
			"DIGIT":  corePkg,
			"HEXDIG": corePkg,
		},
		7159: {
			"DIGIT":  corePkg,
			"HEXDIG": corePkg,
		},
	}
)
