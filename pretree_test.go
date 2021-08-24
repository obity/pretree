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

package pretree_test

import (
	"testing"

	"github.com/obity/pretree"
)

func Test_Match(t *testing.T) {
	// 测试数据data包括 http请求方法，路由规则，客户端请求路径
	data := [][]string{
		{"POST", "/pet/{petId}/uploadImage", "/pet/12121/uploadImage"},
		{"POST", "/pet", "/pet"},
		{"PUT", "/pet", "/pet"},
		{"GET", "/pet/findByStatus", "/pet/findByStatus"},
		{"GET", "/pet/{petId}", "/pet/113"},
		{"GET", "/pet/{petId}/info", "/pet/12121/info"},
		{"POST", "/pet/{petId}", "/pet/12121"},
		{"DELETE", "/pet/{petId}", "/pet/12121"},
		{"GET", "/store/inventory", "/store/inventory"},
		{"POST", "/store/order", "/store/order"},
		{"GET", "/store/order/{orderId}", "/store/order/939"},
		{"DELETE", "/store/order/{orderId}", "/store/order/939"},
		{"POST", "/user/createWithList", "/user/createWithList"},
		{"GET", "/user/{username}", "/user/1002"},
		{"PUT", "/user/{username}", "/user/1002"},
		{"DELETE", "/user/{username}", "/user/1002"},
		{"GET", "/user/login", "/user/login"},
		{"GET", "/user/logout", "/user/logout"},
		{"POST", "/user/createWithArray", "/user/createWithArray"},
		{"POST", "/user", "/user"},
	}

	p := pretree.NewPreTree()
	for _, v := range data {
		method := v[0]
		sourceRule := v[1]
		p.Store(method, sourceRule)
	}

	for _, v := range data {
		method := v[0]
		urlPath := v[2]
		sourceRule := v[1]
		ok, rule, vars := p.Query(method, urlPath)
		if ok && rule == sourceRule {
			t.Logf("urlPath:%s match rule:%s result: %t vars: %s", urlPath, rule, ok, vars)
		} else {
			t.Errorf("method: %s urlPath:%s match rule:%s result: %t", method, urlPath, sourceRule, ok)
		}
	}
}
