// Copyright 2018 The go-freenas AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package integration

import (
	"os"

	"github.com/fishman/go-freenas/freenas"
)

var (
	client *freenas.Client
)

func init() {
	user := os.Getenv("FREENAS_USER")
	password := os.Getenv("FREENAS_PASSWORD")
	server := os.Getenv("FREENAS_SERVER")
	if server == "" || user == "" || password == "" {
		print("!!! Server information is missing. Some tests won't run. !!!\n\n")
		client = freenas.NewClient(server, user, password)
	}
}
