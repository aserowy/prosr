package compiler

// Ast represents the whole body of the parsed tree
type Ast struct {
	nodes []Node
}

// NewAst ctor for Ast
func NewAst() *Ast {
	return new(Ast)
}

// Push adds node to toplevel
func (a *Ast) Push(n Node) {
	a.nodes = append(a.nodes, n)
}

// Fabricate finishes the completly parsed ast
func (a *Ast) Fabricate() []Node {
	// TODO: Validate Ast
	a.pack()

	return a.nodes
}

func (a *Ast) pack() {
	ns := []Node{}
	cp := new(Package)

	for _, n := range a.nodes {
		switch cn := n.(type) {
		case *Package:
			cp = cn
			ns = append(ns, n)
		default:
			if cp == nil {
				ns = append(ns, n)
			} else {
				cp.Nodes = append(cp.Nodes, n)
			}
		}
	}

	a.nodes = ns
}
