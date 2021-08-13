package main

import (
	"fmt"

	"github.com/obity/pretree"
)

func main() {
	p := pretree.NewPreTree()
	p.Store(pretree.MethodGet, "account/{id}/info/:name")
	p.Store(pretree.MethodGet, "account/:id/login")
	p.Store(pretree.MethodGet, "account/{id}")
	p.Store(pretree.MethodGet, "bacteria/count_number_by_month")

	list := []string{"account/929239",
		"account/9929s/login",
		"account/safsd32/info/121323",
		"bacteria/count_number_by_mont",
	}
	for _, v := range list {
		ok, rule, vars := p.Query(pretree.MethodGet, v)
		if ok {
			fmt.Printf("rule: %s url: %s result:%v vars: %v\n", rule, v, ok, vars)

		} else {
			fmt.Printf("rule: %s url: %s result:%v vars: %v\n", "", v, ok, vars)
		}
	}
}
