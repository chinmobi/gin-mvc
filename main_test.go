// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// Do nothing, just as a test skeleton used for real tests.

	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
	}

	for _, test := range tests {
		got := doMain(test.input)

		if got != test.want {
			t.Errorf("got: [%q], want: [%q]", got, test.want)
		}
	}
}

func doMain(input string) string {
	return input
}
