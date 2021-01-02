// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func DemoCase00(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
	}

	assert := assert.New(t)

	for _, test := range tests {
		got := test.input

		assert.Equal(test.want, got)
	}
}
