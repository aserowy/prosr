package compiler

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Builder creates files for the specified language
type Builder struct {
	options map[string]string
}

// NewBuilder ctor for Builder
func NewBuilder(options map[string]string) *Builder {
	b := new(Builder)
	b.options = options

	return b
}

// Build creates files by Ast
func (b *Builder) Build(ast *Ast) {
	fmt.Println(executeParse(buildTemplate(), ast.Nodes))
}

func buildTemplate() *template.Template {
	fm := template.FuncMap{
		"capitalizeFirstLetter": capitalizeFirstLetter,
		"unifyReturnings":       unifyReturnings,
	}

	tmpl, err := template.
		New("charp.tmpl").
		Funcs(fm).
		ParseFiles("compiler/charp.tmpl")

	if err != nil {
		panic(err)
	}

	return tmpl
}

func executeParse(tmpl *template.Template, data interface{}) string {
	b := new(bytes.Buffer)
	err := tmpl.Execute(b, data)
	if err != nil {
		panic(err)
	}

	return b.String()
}

func capitalizeFirstLetter(v string) string {
	switch len(v) {
	case 0:
		return v
	case 1:
		return strings.ToUpper(string(v[0]))
	default:
		return strings.ToUpper(string(v[0])) + strings.ToLower(v[1:])
	}
}

func unifyReturnings(nodes []Node) []Returning {
	ns := map[string]Returning{}
	for _, n := range nodes {
		re := resolveReturning(n)
		if re == nil {
			continue
		}

		ns[re.String()] = (*re)
	}

	r := []Returning{}
	for _, n := range ns {
		r = append(r, n)
	}

	return r
}

func resolveReturning(node Node) *Returning {
	switch n := node.(type) {
	case *Sending:
		if n.Returns == nil {
			return nil
		}

		return resolveReturning(n.Returns)
	case *Returning:
		return n
	default:
		return nil
	}
}
