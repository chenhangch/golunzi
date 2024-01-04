package main

// import (
// 	"bytes"
// 	"flag"
// 	"fmt"
// 	"go/ast"
// 	"go/types"
// 	"log"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// var (
// 	typeNames  = flag.String("type", "", "comma-separated list of type names; must be set")
// 	output     = flag.String("output", "", "output file name; default srcdir/<type>_string.go")
// 	trimprefix = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
// 	buildTags  = flag.String("tags", "", "comma-separated list of build tags to apply")
// 	doc        = flag.Bool("doc", false, "if true only generate error code documentation in markdown format")
// )

// // Usage is a replacement usage function for the flags package.
// // 使用 flag.Usage = Usage 设置命令行参数的使用说明文本
// func Usage() {
// 	fmt.Fprintf(os.Stderr, "Usage of codegen:\n")
// 	fmt.Fprintf(os.Stderr, "\tcodegen [flags] -type T [directory]\n")
// 	fmt.Fprintf(os.Stderr, "\tcodegen [flags] -type T files... # Must be a single package\n")
// 	fmt.Fprintf(os.Stderr, "Flags:\n")
// 	flag.PrintDefaults()
// }

// func main()  {
// 	log.SetFlags(0)
// 	log.SetPrefix("codeg: ")

// 	flag.Usage = Usage
// 	// 命令行解析为定义的标志
// 	flag.Parse()
// 	if len(*typeNames) == 0 {
// 		flag.Usage()
// 		os.Exit(2)
// 	}
// 	types := strings.Split(*typeNames,",")
	
// 	var tags []string
// 	if len(*buildTags) > 0 {
// 		tags = strings.Split(*buildTags,",")
// 	}

// 	// We accept either one directory or a list of files. Which do we have?
// 	args := flag.Args()
// 	if len(args) == 0 {
// 		// Default: process whole package in current directory.
// 		args = []string{"."}
// 	}

// 	// Parse the package once
// 	var dir string
// 	g := Generator{
// 		trimPrefix: *trimprefix,
// 	}


// 	log.Panicf("%s %s",types[1],tags[1])
// }

// // isDir 判断命令文件name是不是一个目录
// func isDir(name string) bool {
// 	info , err := os.Stat(name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return info.IsDir()
// }

// // Generator 生成器
// type Generator struct {
// 	buf bytes.Buffer
// 	pkg *Package

// 	trimPrefix string
// }

// // Printf 输出到&g.buf
// func (g *Generator) Printf(format string, args ...interface{}) {
// 	fmt.Fprintf(&g.buf, format, args...)
// }

// type File struct {
// 	pkg *Package
// 	file *ast.File
// 	typeName string
// 	value []Value
// }

// type Package struct {
// 	name string
// 	defs map[*ast.Ident]types.Object
// 	files []*File
// }



// type Value struct {
// 	originalName string
// 	name string


// 	value uint64
// 	signed bool
// 	str string
// }


// func ParseComment(varName, comment string) (httpCode, desc string) {
// 	reg := regexp.MustCompile(`\w\s*-\s*(\d{3})\s*:\s*([A-Z].*)\s*\.\n*`)
// 	if !reg.MatchString(comment) {
// 		return "500", "Internal server error"
// 	}
// 	groups := reg.FindStringSubmatch(comment)
// 	//隐含有3块分组
// 	if len(groups) != 3 {
// 		return "500", "Internal server error"
// 	}
// 	return groups[1], groups[2]
// }
