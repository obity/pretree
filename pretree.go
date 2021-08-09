/*
   Copyright (c) 2021 ffactory.org
   hotbuild is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/
package pretree

import (
	"strings"
)

const (
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)

type HttpMethod string

var treeMap map[HttpMethod]*Tree

func init() {
	methods := []HttpMethod{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}
	treeMap = make(map[HttpMethod]*Tree)
	for _, method := range methods {
		tree := newTree()
		tree.Name = string(method)
		treeMap[method] = tree
	}
}

type Tree struct {
	Name       string
	Nodes      []*Tree
	isEnd      bool
	isVariable bool
}

func newTree() *Tree {
	return &Tree{}
}

func GetTreeByMethod(method HttpMethod) *Tree {
	return treeMap[method]
}

func (t *Tree) Insert(url string) {
	root := t
	list := parsePath(url)
	for _, word := range list {
		isExist := false
		// 如果已经存在路径，继续匹配子节点
		for _, n := range root.Nodes {
			if n.Name == word {
				isExist = true
				root = n
				break
			}
		}
		// 已存在进入下一次循环
		if isExist {
			continue
		}
		// 不存在的路径新增
		node := newTree()
		node.Name = word
		// 记录本路径是否变量
		if isVariable(word) {
			node.isVariable = true
		}
		root.Nodes = append(root.Nodes, node)
		root = node
	}
	root.isEnd = true
}

func (t *Tree) Match(url string) (bool, *Tree) {
	root := t
	list := parsePath(url)
	for index, word := range list {
		isExist := false
		preVar := false
		for _, n := range root.Nodes {
			if n.Name == word {
				preVar = false
				isExist = true
				root = n
				break
			}
		}
		if isExist {
			continue
		}
		// 第二个路径匹配不到情况下，查找是否有变量路径，继续从变量路径往下找
		for _, m := range root.Nodes {
			if m.isVariable && index > 0 && !preVar {
				preVar = true
				root = m
				break
			} else {
				return false, nil
			}
		}
	}
	if root.isEnd {
		return true, root
	} else {
		return false, nil
	}
}

func (t *Tree) Next() []*Tree {
	return t.Nodes
}

func parsePath(path string) []string {
	return strings.Split(path, "/")
}

func isVariable(s string) bool {
	return strings.HasPrefix(s, ":")
}
