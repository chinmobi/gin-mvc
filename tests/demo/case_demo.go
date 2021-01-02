// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package demo

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chinmobi/gin-mvc/tests"

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

func DemoCase01(t *testing.T) {
	router := tests.GinEngine()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	log.Println(w.Body.String())
}
