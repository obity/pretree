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

package pretree

import (
	"strings"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type PreTree struct {
	treeGroup map[string]*Tree
}

// 初始化对象
//
// Initialize object
func NewPreTree() *PreTree {
	p := &PreTree{}
	methods := []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}
	p.treeGroup = make(map[string]*Tree)
	for _, method := range methods {
		tree := newTree()
		tree.name = string(method)
		p.treeGroup[method] = tree
	}
	return p
}

// 存储路由规则
//
// Store routing rules
func (p *PreTree) Store(method, urlRule string) {
	t := p.treeGroup[method]
	t.insert(urlRule)
}

// 查询URL匹配的树节点并返回变量
//
// Query the tree node with matching URL and return variables
func (p *PreTree) Query(method, urlPath string) (isExist bool, rule string, vars map[string]string) {
	t := p.treeGroup[method]
	isExist, node, vars := t.match(urlPath)
	if isExist {
		return true, node.Rule(), vars
	} else {
		return false, "", vars
	}
}

// 前缀树数据结构
//
// Prefix tree data structure
type Tree struct {
	rule       string
	name       string
	nodes      []*Tree
	isEnd      bool
	isVariable bool
}

func newTree() *Tree {
	return &Tree{}
}

func (t *Tree) appendChild(child *Tree) {
	t.nodes = append(t.nodes, child)
}

// 获取子节点
//
// Get child nodes
func (t *Tree) Child() []*Tree {
	return t.nodes
}

// 获取当前节点的路由规则
//
// Get the routing rule of the current node
func (t *Tree) Rule() string {
	return t.rule
}

// 获取当前节点的名称
//
// Get the name of the current node
func (t *Tree) Name() string {
	return t.name
}

// 获取当前节点的变量名
//
// Get the variable name of the current node
//	:id => id
func (t *Tree) VarName() string {
	return strings.TrimPrefix(t.name, ":")
}

func (t *Tree) insert(urlRule string) {
	current := t
	list := parsePath(urlRule)
	for _, word := range list {
		isExist := false
		// 如果已经存在路径，继续匹配子节点
		for _, n := range current.Child() {
			if n.name == word {
				isExist = true
				current = n
				break
			}
		}
		// 已存在进入下一次循环
		if isExist {
			continue
		}
		// 不存在的路径新增
		node := newTree()
		node.name = word
		// 记录本路径是否变量
		if isVariable(word) {
			node.isVariable = true
		}
		current.appendChild(node)
		current = node
	}
	current.rule = urlRule
	current.isEnd = true
}

func (t *Tree) match(urlPath string) (bool, *Tree, map[string]string) {
	// vars 用于存储路由变量数据
	vars := make(map[string]string)
	current := t
	list := parsePath(urlPath)
	for index, word := range list {
		isExist := false
		hasVar := false
		for _, n := range current.Child() {
			if n.name == word {
				hasVar = false
				isExist = true
				current = n
				break
			}
		}
		if isExist {
			continue
		}
		// 第二个路径匹配不到情况下，查找是否有变量路径，继续从变量路径往下找
		for _, m := range current.Child() {
			if m.isVariable && index > 0 && !hasVar {
				hasVar = true
				current = m
				vars[m.VarName()] = word
				break
			}
		}
		// 找到有变量路径,进入下一次循环
		if hasVar {
			continue
		}
	}
	if current.isEnd {
		return true, current, vars
	} else {
		return false, nil, nil
	}
}

func parsePath(path string) []string {
	path = formatRule(path)
	return strings.Split(path, "/")
}

func formatRule(rule string) string {
	rule = strings.ReplaceAll(rule, "{", ":")
	rule = strings.ReplaceAll(rule, "}", "")
	rule = strings.TrimPrefix(rule, "/")
	rule = strings.TrimSuffix(rule, "/")
	return rule
}

func isVariable(s string) bool {
	return strings.HasPrefix(s, ":")
}
