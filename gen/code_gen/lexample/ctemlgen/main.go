package main

import (
	"os"
	"text/template"
)

type Field struct {
    Name string
    Type string
}

type Struct struct {
    Name   string
    Fields []Field
}

func main() {
    // 定义要传递给模板的参数
    data := struct {
        PackageName string
        Structs     []Struct
    }{
        PackageName: "example",
        Structs: []Struct{
            {
                Name: "User",
                Fields: []Field{
                    {Name: "ID", Type: "int"},
                    {Name: "Name", Type: "string"},
                    {Name: "Age", Type: "int"},
                },
            },
            {
                Name: "Product",
                Fields: []Field{
                    {Name: "ID", Type: "int"},
                    {Name: "Name", Type: "string"},
                    {Name: "Price", Type: "float64"},
                },
            },
        },
    }

    // 解析模板文件
    tmpl, err := template.ParseFiles("./template.tmpl")
    if err != nil {
        panic(err)
    }

    // 执行模板并输出到文件
    f, err := os.Create("./out/output.go")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    err = tmpl.Execute(f, data)
    if err != nil {
        panic(err)
    }
}