// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"math/big"

	"github.com/ogen-go/ogen/ogenregex"
)

var regexMap = map[string]ogenregex.Regexp{
	"^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$": ogenregex.MustCompile("^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$"),
	"^\\d-\\d$": ogenregex.MustCompile("^\\d-\\d$"),
	"foo.*":     ogenregex.MustCompile("foo.*"),
	"string_.*": ogenregex.MustCompile("string_.*"),
}
var ratMap = map[string]*big.Rat{
	"10": func() *big.Rat {
		r, ok := new(big.Rat).SetString("10")
		if !ok {
			panic(fmt.Sprintf("rat %q: can't parse", "10"))
		}
		return r
	}(),
	"5": func() *big.Rat {
		r, ok := new(big.Rat).SetString("5")
		if !ok {
			panic(fmt.Sprintf("rat %q: can't parse", "5"))
		}
		return r
	}(),
}
