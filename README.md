# Pretree

[![GitHub release](https://img.shields.io/github/v/release/obity/pretree.svg?color=peru)](https://github.com/obity/pretree/releases/latest)
[![Released API docs](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/obity/pretree)
[![Build](https://img.shields.io/github/actions/workflow/status/obity/pretree/.github/workflows/go.yml?branch=master)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/obity/pretree?color=pink)](https://goreportcard.com/report/github.com/obity/pretree)
[![Lines of code](https://img.shields.io/tokei/lines/github/obity/pretree.svg?color=beige)](#)
[![Languages](https://img.shields.io/github/languages/top/obity/pretree.svg?color=yellow)](#)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/obity/pretree)](#)
[![GPL3 licensed](https://img.shields.io/github/license/obity/pretree.svg)](./LICENSE)

pretree is a package for storing and querying routing rules with prefix tree .

pretree 是一个用于存储和查询路由规则的包。它用前缀树存储路由规则，支持包含变量的路由。

pretree is a package for storing and querying routing rules. It uses prefix tree to store routing rules and supports routing with variables.

# Doc

See this document at [GoDoc](https://pkg.go.dev/github.com/obity/pretree)

# Install
    
    go get -u github.com/obity/pretree@latest
    
# ToDo

- [x] Query返回路由中的变量vars存在一个map中.
