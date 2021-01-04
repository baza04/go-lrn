package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var firstTestCase = `1
2
3
4
4
5
5`

var firstCaseRes = `1
2
3
4
4
5
5
`

func testingOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(firstTestCase))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("first test failed")
	}

	result := out.String()
	if result != firstCaseRes {
		t.Errorf("first test failed - result not match\n res: %v\nexp: %v", result, firstCaseRes)
	}
}
