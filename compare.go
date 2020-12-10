/*
Copyright © 2020 Neil Hemming

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package testsupport provides support functions for testing
package testsupport

import (
	"strings"
	"testing"
)

//CompareStrings compares 2 strings for differences
func CompareStrings(t *testing.T, expected string, got string) {

	spaces := strings.Repeat(" ", 4)
	expected = strings.Replace(strings.Replace(expected, "\t", spaces, -1), "\r\n", "\n", -1)
	got = strings.Replace(strings.Replace(got, "\t", spaces, -1), "\r\n", "\n", -1)

	eRune := []rune(expected)
	gRune := []rune(got)

	eLen := len(eRune)
	gLen := len(gRune)

	var min int
	if eLen > gLen {
		min = gLen
	} else {
		min = eLen
	}

	mismatch := false

	if eLen != gLen {
		t.Errorf("Length mismatch:%v %v", eLen, gLen)
		mismatch = true
	}

	// compare rune by rune
	for i := 0; i < min; i++ {
		if eRune[i] == gRune[i] {
			continue
		}

		// break
		t.Errorf("Mismatch %d:\nexp:%v\ngot:%v\nRune %d(%v) %d(%v)", i,
			expected, string(gRune[:i+1]),
			int(eRune[i]), string(eRune[i]), int(gRune[i]), string(gRune[i]))

		mismatch = true

		break
	}

	if mismatch {
		t.Errorf("Got:\n%v\nExpected:\n%v", got, expected)
	}
}
