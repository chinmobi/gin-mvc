// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// at https://github.com/gin-gonic/gin/blob/master/LICENSE

package internal

import (
	"bytes"
	"path"

	"github.com/chinmobi/gin-mvc/evt/internal/bytesconv"
)

var (
	strColon = []byte(":")
	strStar  = []byte("*")
)

func countParams(path string) uint16 {
	var n uint16
	s := bytesconv.StringToBytes(path)
	n += uint16(bytes.Count(s, strColon))
	n += uint16(bytes.Count(s, strStar))
	return n
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func assert1(guard bool, text string) {
	if !guard {
		panic(text)
	}
}
