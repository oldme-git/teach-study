package basetest

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"testing"
)

func TestReplaceFunc(t *testing.T) {
	var (
		patternStr = `([A-Z])\w+`
		str        = "hello Golang 2018~2021!"
	)
	// In contrast to [ExampleReplaceFunc]
	// the result contains the `pattern' of all subpatterns that use the matching function
	result, err := gregex.ReplaceFuncMatch(patternStr, []byte(str), func(match [][]byte) []byte {
		g.Dump(match)
		return []byte("Gf")
	})
	g.Dump(result)
	g.Dump(err)

	// ReplaceStringFuncMatch
	resultStr, err := gregex.ReplaceStringFuncMatch(patternStr, str, func(match []string) string {
		g.Dump(match)
		match[0] = "Gf"
		return match[0]
	})
	g.Dump(resultStr)
	g.Dump(err)
}

// [
// 	"Golang",
// 	"G",
// ]
// "hello Gf 2018~2021!"
// <nil>
// [
// 	"Golang",
// 	"G",
// ]
// "hello Gf 2018~2021!"
// <nil>
