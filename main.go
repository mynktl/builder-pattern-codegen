package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"strings"
)

var builder = []string{BuildTemplate, BuildVarTemplate, ValidateTemplate, BuildFuncTemplate}
var predicate = []string{PredicateTemplate, PredicateVarTemplate}
var utils = []string{SetVarTemplate, GetVarTemplate}

const (
	Builder = iota
	Predicate
	Utils
	All
)

type generator struct {
	file       *string
	genType    int
	dir        *string
	headerFile *string
}

func newGenerator(file, dir, headerFile *string, gentype int) *generator {
	return &generator{
		file:       file,
		genType:    gentype,
		dir:        dir,
		headerFile: headerFile,
	}
}

func (g *generator) newParser() *ast.File {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, *g.file, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return f
}

func (g *generator) copyHeaderText(df *os.File) {
	sf, err := os.Open(*g.headerFile)
	if err != nil {
		log.Fatalf("Failed to open header file : %v", err)
	}

	if _, err := io.Copy(df, sf); err != nil {
		log.Fatalf("Failed to copy header file : %v", err)
	}
	sf.Close()
}

func (g *generator) generatePackage(template []string, file string) {
	fs := g.newParser()
	if fs == nil {
		log.Fatalf("Error occurred creating a new parser")
	}

	f, err := os.Create(*g.dir + "/" + file)
	if err != nil {
		log.Fatalf("Failed to create new file : %v", err)
	}

	g.copyHeaderText(f)

	for _, t := range template {
		if err := parseTemplate(t, fs, bufio.NewWriter(f)); err != nil {
			log.Fatalf("Failed to parse template : %v", err)
		}
	}
	err = f.Close()
	if err != nil {
		log.Fatalf("Failed to close a file : %v", err)
	}
	return
}

func (g *generator) generateBuilder() {
	g.generatePackage(builder, "builder.go")
}

func (g *generator) generatePredicate() {
	g.generatePackage(predicate, "predicate.go")
}

func (g *generator) generateUtils() {
	g.generatePackage(utils, "utils.go")
}

func (g *generator) Execute() {
	switch g.genType {
	case Builder:
		g.generateBuilder()
	case Predicate:
		g.generatePredicate()
	case Utils:
		g.generateUtils()
	case All:
		for i := Builder; i < All; i++ {
			g.genType = i
			g.Execute()
		}
	}
}

func main() {
	dir, _ := os.Getwd()

	filename := flag.String("file", "", "file name")
	outputpath := flag.String("dir", dir, "output directory name")
	bfile := flag.String("boilerplate", "", "boilerplate file path")

	flag.Parse()

	if len(*filename) == 0 || len(*bfile) == 0 {
		log.Fatalf("File name is missing")
	}
	generator := newGenerator(filename, outputpath, bfile, All)

	generator.Execute()
	return
}

func parseTemplate(template string, fs *ast.File, w *bufio.Writer) error {
	var objname string
	shouldParseVar := true
	predicateGenertor := false

	if template == BuildTemplate || template == PredicateTemplate || template == BuildFuncTemplate || template == ValidateTemplate {
		shouldParseVar = false
	}
	if template == PredicateVarTemplate {
		predicateGenertor = true
	}

	// fs is a parsed, type-checked *ast.File.
	ast.Inspect(fs, func(n ast.Node) bool {
		switch b := n.(type) {
		case *ast.GenDecl:
			if b.Tok != token.TYPE {
				break
			}
			objname = b.Specs[0].(*ast.TypeSpec).Name.Name
		}
		return true
	})

	// copy strcture if it is builder template
	if template == BuildTemplate {
		copyStructFile(fs, w)
	}

	template = strings.Replace(template, "$NewObj", objname, -1)
	template = strings.Replace(template, "$newobj", objname, -1)
	template = strings.Replace(template, "$obj", string(strings.ToLower(objname)[0]), -1)
	template = strings.Replace(template, "$Struct", objname, -1)

	// fs is a parsed, type-checked *ast.File.
	ast.Inspect(fs, func(n ast.Node) bool {
		if expr, ok := n.(*ast.StructType); ok && shouldParseVar {
			for _, field := range expr.Fields.List {
				template := strings.Replace(template, "$Var", field.Names[0].Name, -1)
				template = strings.Replace(template, "$var", field.Names[0].Name, -1)
				ntype := fmt.Sprintf("%s", field.Type)
				template = strings.Replace(template, "$iType", ntype, -1)
				if predicateGenertor {
					fcond := ""
					switch ntype {
					case "string":
						fcond = "len(%s.%s) == 0"
					case "bool":
						fcond = "%s.%s == true"
					case "int":
						fcond = "%s.%s == 0"
					}
					cond := fmt.Sprintf(fcond, string(strings.ToLower(objname)[0]), field.Names[0].Name)
					template = strings.Replace(template, "$cond", cond, -1)
				}
				_, err := w.WriteString(template)
				if err != nil {
					log.Fatalf("Failed to write modified template : %v", err)
				}
			}
		}
		return true
	})

	if !shouldParseVar {
		_, err := w.WriteString(template)
		if err != nil {
			log.Fatalf("Failed to write modified template : %v", err)
		}
	}
	if err := w.Flush(); err != nil {
		log.Fatalf("Failed to flush data to file : %v", err)
	}
	return nil
}

func copyStructFile(fs *ast.File, w *bufio.Writer) {
	// import error package
	_, err := fmt.Fprintf(w, "\nimport \"github.com/pkg/errors\"\n")
	if err != nil {
		log.Fatalf("Failed to write struct : %v", err)
	}

	// fs is a parsed, type-checked *ast.File.
	ast.Inspect(fs, func(n ast.Node) bool {
		switch b := n.(type) {
		case *ast.GenDecl:
			if b.Tok != token.TYPE {
				break
			}
			_, err := fmt.Fprintf(w, "\ntype %s struct {\n", b.Specs[0].(*ast.TypeSpec).Name.Name)
			if err != nil {
				log.Fatalf("Failed to write struct : %v", err)
			}
		}
		return true
	})

	// fs is a parsed, type-checked *ast.File.
	ast.Inspect(fs, func(n ast.Node) bool {
		if expr, ok := n.(*ast.StructType); ok {
			for _, field := range expr.Fields.List {
				_, err := fmt.Fprintf(w, "\t%s\n", field.Comment.List[0].Text)
				if err != nil {
					log.Fatalf("Failed to write struct : %v", err)
				}
				_, err = fmt.Fprintf(w, "\t%s %s\n", field.Names[0].Name, field.Type)
				if err != nil {
					log.Fatalf("Failed to write struct : %v", err)
				}
			}
		}
		return true
	})

	// Add predicatelist variable
	if _, err := w.WriteString(BuilderFieldTemplate); err != nil {
		log.Fatalf("Failed to add predicate and error filed : %v", err)
	}

	_, err = fmt.Fprintf(w, "}\n")
	if err != nil {
		log.Fatalf("Failed to write struct : %v", err)
	}
	if err := w.Flush(); err != nil {
		log.Fatalf("Failed to flush data to file : %v", err)
	}
}
