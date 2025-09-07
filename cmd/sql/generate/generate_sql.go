package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"strings"
)

const query = "query"

func main() {
	fset := token.NewFileSet()
	node := &ast.File{
		Name: ast.NewIdent("sql"),
	}

	dirEntrys, err := os.ReadDir(query)
	if err != nil {
		fmt.Printf("Read dir: %s", err.Error())
		os.Exit(1)
	}

	node.Decls = make([]ast.Decl, len(dirEntrys)+1)

	variables := make([]string, len(dirEntrys))
	for i, dirEntry := range dirEntrys {
		for _, word := range strings.Split(dirEntry.Name(), "_") {
			rIndex := len(word)

			if pointIdex := strings.Index(word, "."); pointIdex != -1 {
				rIndex = pointIdex
			}

			variables[i] = variables[i] + strings.ToUpper(string(word[0])) + word[1:rIndex]
		}
	}

	currentDecls := 0
	node.Decls[currentDecls] = &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Name: &ast.Ident{
					Name: "_",
				},
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"embed\"",
				},
			},
		},
	}

	for i, variable := range variables {
		currentDecls++
		node.Decls[currentDecls] = &ast.GenDecl{
			Doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: fmt.Sprintf("//go:embed %s/%s", query, dirEntrys[i].Name()),
					},
				},
			},
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						{
							Name: variable,
						},
					},
					Type: &ast.Ident{
						Name: "string",
					},
				},
			},
		}
	}

	f, err := os.Create(fmt.Sprintf("%s.go", query))
	if err != nil {
		fmt.Printf("Create file: %s\n", err.Error())
		os.Exit(1)
	}

	format.Node(f, fset, node)
}
