package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goparse"
	app.Usage = "print list of all top-level interfaces defined in a go source file"
	app.UsageText = "goiface filename"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		{
			Name:  "Nathan Jordan",
			Email: "nathan.m.jordan@gmail.com",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 1 {
			return cli.ShowAppHelp(c)
		}
		filename := c.Args().First()
		interfaces, err := getInterfacesFromFilename(filename)
		if err != nil {
			return err
		}
		for _, name := range interfaces {
			fmt.Println(name)
		}
		return nil
	}

	app.Run(os.Args)
}

func getInterfacesFromFilename(filename string) ([]string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}
	return parseDeclarations(f), nil
}

func parseDeclarations(file *ast.File) []string {
	var interfaces []string
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if found, name := getInterfaceName(genDecl); found {
			interfaces = append(interfaces, name)
		}
	}
	return interfaces
}

func getInterfaceName(decl *ast.GenDecl) (found bool, name string) {
	for _, spec := range decl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			return false, ""
		}

		if _, ok := typeSpec.Type.(*ast.InterfaceType); !ok {
			return false, ""
		}
		return true, typeSpec.Name.Name
	}
	return false, ""
}
