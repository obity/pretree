/*
   Copyright (c) 2021 ffactory.org
   pretree is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/
package main

import (
	"fmt"

	"github.com/obity/pretree"
)

func main() {
	pretree.Store(pretree.MethodGet, "account/{id}/info/:2333")
	pretree.Store(pretree.MethodGet, "account/:id/login")
	pretree.Store(pretree.MethodGet, "account/{id}")
	pretree.Store(pretree.MethodGet, "bacteria/count_number_by_month")

	list := []string{"account/929239",
		"account/9929s/login",
		"account/safsd32/info/121323",
		"bacteria/count_number_by_mont",
	}
	for _, v := range list {
		ok, tree := pretree.Query(pretree.MethodGet, v)
		if ok {
			fmt.Printf("rule: %s url: %s result:%v\n", tree.Rule(), v, ok)

		} else {
			fmt.Printf("rule: %s url: %s result:%v\n", "", v, ok)
		}
	}
}
