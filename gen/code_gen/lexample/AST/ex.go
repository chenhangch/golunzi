package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

func main() {
	// 1. 解析Go源代码文件，生成AST抽象语法树。
	fset := token.NewFileSet()

	absPath, _ := filepath.Abs("./code.go")

	astFile, err := parser.ParseFile(fset, absPath, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 遍历AST树
	ast.Inspect(astFile, func(n ast.Node) bool {
		// 判断是否为注解
		if comment, ok := n.(*ast.CommentGroup); ok {
		// 输出注解内容
			fmt.Println(comment.Text())
		}
		return true
	})
}