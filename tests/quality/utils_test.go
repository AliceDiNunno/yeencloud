package quality

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

func PackageVariables(pkgPath string) ([]*ast.ValueSpec, error) {
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, pkgPath, nil, 0)
	if err != nil {
		log.Fatal("Failed to parse package:", err)
		return nil, err
	}

	var variables []*ast.ValueSpec
	for _, pack := range packs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.GenDecl); isFn {
					for _, s := range fn.Specs {
						if v, isVal := s.(*ast.ValueSpec); isVal {
							variables = append(variables, v)
						}
					}
				}
			}
		}
	}
	return variables, nil
}

func getVariables(pkg string) ([]variable, error) {
	vars, err := PackageVariables(pkg)
	if err != nil {
		return nil, err
	}

	var variables []variable
	for _, v := range vars {
		composite, ok := v.Values[0].(*ast.CompositeLit)
		if ok {
			var variable variable

			variable.Name = v.Names[0].Name
			ident, ok := composite.Type.(*ast.Ident)
			if ok {
				variable.Type = ident.Name
				for _, elt := range composite.Elts {
					kv, ok := elt.(*ast.KeyValueExpr)
					if !ok {
						continue
					}

					if variable.data == nil {
						variable.data = make(map[string]interface{})
					}

					variable.data["fields"] = kv
				}
			}

			variables = append(variables, variable)
		}
	}

	return variables, nil
}

func listLanguageFiles(localizationPath string) []string {
	files, err := os.ReadDir(localizationPath)
	if err != nil {
		log.Fatalln("Error reading translation directory " + err.Error())
	}

	var languageFiles []string
	for _, file := range files {
		path := fmt.Sprintf("%s/%s", localizationPath, file.Name())

		if strings.HasSuffix(path, ".toml") {
			languageFiles = append(languageFiles, path)
		}
	}
	return languageFiles
}

func loadTomlFile(path string) (map[string]interface{}, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	_, err = toml.Decode(string(file), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
