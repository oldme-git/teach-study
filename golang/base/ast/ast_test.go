package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func getDemoGoPath() string {
	demoGo, _ := os.Getwd()
	demoGo = filepath.FromSlash(demoGo + "/demo.go")
	return demoGo
}

// 获取token
func TestScanner(t *testing.T) {
	demoGo := getDemoGoPath()

	demoFile, err := os.Open(demoGo)
	if err != nil {
		panic(err)
	}
	defer demoFile.Close()

	fStat, err := demoFile.Stat()
	if err != nil {
		panic(err)
	}
	size := fStat.Size()

	src := make([]byte, size)
	_, err = demoFile.Read(src)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	file := fset.AddFile(demoGo, fset.Base(), int(size))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}

func TestScanner2(t *testing.T) {
	src := []byte(`
func demo() {
	fmt.Println("When you are old and gray and full of sleep")
}
`)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}

// 获取AST
func TestParse(t *testing.T) {
	demoGo := getDemoGoPath()
	// 创建一个文件集
	fset := token.NewFileSet()
	// 解析文件，返回一个 *ast.File 类型的值，表示解析后的 AST
	f, err := parser.ParseFile(fset, demoGo, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
}

func TestPrintAST(t *testing.T) {
	src := `
package main
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)
}

func ExampleGetStructName() {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, getDemoGoPath(), nil, parser.ParseComments)
	if err != nil {
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if v, ok := n.(*ast.TypeSpec); ok {
			fmt.Println(v.Name.Name)
		}
		return true
	})

	// Output:
	// A
	// B
	// C
}
