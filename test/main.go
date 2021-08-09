package main

import (
	"fmt"

	"github.com/obity/pretree"
)

func main() {
	getTree := pretree.GetTreeByMethod(pretree.HttpMethod("GET"))
	getTree.Insert("account/{id}/info/:2333")
	getTree.Insert("account/:id/login")
	getTree.Insert("account/{id}")
	getTree.Insert("bacteria/count_number_by_month")

	list := []string{"account/929239", "account/9929s/login", "account/safsd32/info/121323", "bacteria/count_number_by_month"}
	for _, v := range list {
		ok, _ := getTree.Match(v)
		fmt.Printf("%s:%t\n", v, ok)
	}
}
