package main

import (
	"flag"
	"fmt"
	"github.com/elimity-com/abnf"
	"io/ioutil"
	"log"
)

func main() {
	input := flag.String("i", "", "source file of the abnf syntax")
	output := flag.String("o", "", "destination for the generated abnf parser")
	flag.Parse()

	raw, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	f := abnf.GenerateABNFAsAlternatives("syntax", string(raw))
	if err := ioutil.WriteFile(*output, []byte(fmt.Sprintf("%#v", f)), 0644); err != nil {
		log.Fatal(err)
	}
}
