package main

import (
	"fmt"

	"github.com/obity/pretree"
)

func main() {
	pretree.Store(pretree.MethodGet, "account/{id}/info/:name")
	pretree.Store(pretree.MethodGet, "account/:id/login")
	pretree.Store(pretree.MethodGet, "account/{id}")
	pretree.Store(pretree.MethodGet, "bacteria/count_number_by_month")

	list := []string{"account/929239",
		"account/9929s/login",
		"account/safsd32/info/121323",
		"bacteria/count_number_by_mont",
	}
	for _, v := range list {
		ok, tree, vars := pretree.Query(pretree.MethodGet, v)
		if ok {
			fmt.Printf("rule: %s url: %s result:%v vars: %v\n", tree.Rule(), v, ok, vars)

		} else {
			fmt.Printf("rule: %s url: %s result:%v vars: %v\n", "", v, ok, vars)
		}
	}
}
