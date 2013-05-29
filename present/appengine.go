// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package main

import (
	_ "github.com/syntropy/golang-slides/pkg/playground"
	"github.com/syntropy/golang-slides/pkg/present"
)

var basePath = "./present/"

func init() {
	playScript(basePath, "jquery.js", "playground.js")
	present.PlayEnabled = true
}
