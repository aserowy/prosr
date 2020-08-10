package compiler

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Builder creates files for the specified language
type Builder struct {
	Options map[string]string
	Ast     []Node
}

// NewBuilder ctor for Builder
func NewBuilder(path string, ast *Ast, options map[string]string) *Builder {
	b := new(Builder)
	b.Options = options
	b.Ast = ast.Nodes

	return b
}

// Build creates files by Ast
func (b *Builder) Build() {
	fmt.Println(executeParse(buildTemplate(), b))
}

func buildTemplate() *template.Template {
	tm := map[string]string{}

	fm := template.FuncMap{
		"addType":               func(key string, value string) bool { tm[key] = value; return true },
		"type":                  func(key string) string { return resolveOption(tm, key) },
		"capitalizeFirstLetter": capitalizeFirstLetter,
		"resolveOption":         resolveOption,
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

func resolveOption(options map[string]string, key string) string {
	if v, ok := options[key]; ok {
		return v
	}

	return key
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
