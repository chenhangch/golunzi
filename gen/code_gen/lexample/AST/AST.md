# What is AST
AST，全名是abstract syntax tree（抽象语法树），是一种使用树状结构表示源代码的语法结构，树的每一个节点表示源代码中的一个结构

> 使用中序遍历可以还原表达式

## 2.为什么需要AST
+ 代码语法的检查、代码风格的检查、代码的格式化、代码的高亮、代码错误提示、代码自动补全等等
+ 代码混淆压缩
+ 优化变更代码，改变代码结构

> 由此，我们可以自定义**代码生成器**，这个非常有用

## 3.Go中的AST
go的官方提供了几个包，可以帮助我们进行AST分析
+ go/scanner: 词法解析，将源代码分割成一个个token
+ go/token: token类型以及相关结构体定义
+ go/ast: ast的结构定义
+ go/parse: 语法分析，读取token流生成ast

抽象语法树由节点**Node**构成，Golang的AST由三个节点构成：分别是`表达式和类型节点(Expressions and type nodes)` \ `语句节点(statement nodes)` \ `声明节点(declaration nodes)`

每个节点都包含标识其在源代码中的开头和结尾位置的信息

示例：
```go
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

}
```
运行后输出：
```
     0  *ast.File {
     1  .  Package: 2:1
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: 2:9
     4  .  .  Name: "main"
     5  .  }
     6  .  Decls: []ast.Decl (len = 1) {
     7  .  .  0: *ast.FuncDecl {
     8  .  .  .  Name: *ast.Ident {
     9  .  .  .  .  NamePos: 3:6
    10  .  .  .  .  Name: "main"
    11  .  .  .  .  Obj: *ast.Object {
    12  .  .  .  .  .  Kind: func
    13  .  .  .  .  .  Name: "main"
    14  .  .  .  .  .  Decl: *(obj @ 7)
    15  .  .  .  .  }
    16  .  .  .  }
    17  .  .  .  Type: *ast.FuncType {
    18  .  .  .  .  Func: 3:1
    19  .  .  .  .  Params: *ast.FieldList {
    20  .  .  .  .  .  Opening: 3:10
    21  .  .  .  .  .  Closing: 3:11
    22  .  .  .  .  }
    23  .  .  .  }
    24  .  .  .  Body: *ast.BlockStmt {
    25  .  .  .  .  Lbrace: 3:13
    26  .  .  .  .  List: []ast.Stmt (len = 1) {
    27  .  .  .  .  .  0: *ast.ExprStmt {
    28  .  .  .  .  .  .  X: *ast.CallExpr {
    29  .  .  .  .  .  .  .  Fun: *ast.Ident {
    30  .  .  .  .  .  .  .  .  NamePos: 4:2
    31  .  .  .  .  .  .  .  .  Name: "println"
    32  .  .  .  .  .  .  .  }
    33  .  .  .  .  .  .  .  Lparen: 4:9
    34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
    35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
    36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10
    37  .  .  .  .  .  .  .  .  .  Kind: STRING
    38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
    39  .  .  .  .  .  .  .  .  }
    40  .  .  .  .  .  .  .  }
    41  .  .  .  .  .  .  .  Ellipsis: -
    42  .  .  .  .  .  .  .  Rparen: 4:25
    43  .  .  .  .  .  .  }
    44  .  .  .  .  .  }
    45  .  .  .  .  }
    46  .  .  .  .  Rbrace: 5:1
    47  .  .  .  }
    48  .  .  }
    49  .  }
    50  .  FileStart: 1:1
    51  .  FileEnd: 5:3
    52  .  Scope: *ast.Scope {
    53  .  .  Objects: map[string]*ast.Object (len = 1) {
    54  .  .  .  "main": *(obj @ 11)
    55  .  .  }
    56  .  }
    57  .  Unresolved: []*ast.Ident (len = 1) {
    58  .  .  0: *(obj @ 29)
    59  .  }
    60  }
```
