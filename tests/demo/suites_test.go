// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package demo

import (
	"testing"
)

func runCase(t *testing.T, testCase func(*testing.T)) {
	Before()
	defer After()

	testCase(t)
}

func TestRunSuites(t *testing.T) {
	SetUp()
	defer TearDown()

	runCase(t, DemoCase00)
}
