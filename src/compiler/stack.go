package compiler

// Stack handles parsed token from listener
type Stack struct {
	entries []Node
}

// NewStack ctor for Stack
func NewStack() *Stack {
	return new(Stack)
}

// Filled checks if stack has at least one item
func (ns *Stack) Filled() bool {
	return len(ns.entries) > 0
}

// Push adds entry to the end of stack
func (ns *Stack) Push(n Node) {
	ns.entries = append(ns.entries, n)
}

// Pop returns last item
func (ns *Stack) Pop() Node {
	if ns.Filled() == false {
		panic("stack is empty unable to pop")
	}

	r := ns.entries[len(ns.entries)-1]
	ns.entries = ns.entries[:len(ns.entries)-1]

	return r
}

// Empty removes all entries from stack, returns true if entries were removed
func (ns *Stack) Empty() bool {
	r := false
	if ns.Filled() {
		r = true
	}

	ns.entries = ns.entries[0:0]

	return r
}

// PeekToken checks first item and compares tokens
func (ns *Stack) PeekToken(token string) bool {
	if ns.Filled() == false {
		return false
	}

	return ns.entries[len(ns.entries)-1].TokenLiteral() == token
}
