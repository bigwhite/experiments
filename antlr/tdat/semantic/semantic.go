package semantic

import (
	"errors"
	"fmt"
	"math"
)

type WindowsRange struct {
	low  int
	high int
}

func NewWindowsRange(low, high int) WindowsRange {
	return WindowsRange{
		low:  low,
		high: high,
	}
}

type Model struct {
	// conditionExpr
	t Tree

	// windowsRange
	wr WindowsRange

	// enumerableFunc
	ef string

	// result
	result []string
}

func NewModel(reversePolishExpr []Value, wr WindowsRange, ef string, result []string) *Model {
	m := &Model{
		t:      NewFrom(reversePolishExpr),
		wr:     wr,
		ef:     ef,
		result: result,
	}
	return m
}

func (m *Model) outputResult(in map[string]interface{}) (map[string]interface{}, error) {
	if len(m.result) == 0 {
		// all metric are needed
		return in, nil
	}

	var result = make(map[string]interface{})
	for _, k := range m.result {
		if v, ok := in[k]; ok {
			result[k] = v
		} else {
			return nil, fmt.Errorf("not found metric: %s", k)
		}

	}

	return result, nil
}

var (
	ErrNotMeetNone    = errors.New("not meet none")
	ErrNotMeetEach    = errors.New("not meet each")
	ErrNotMeetAny     = errors.New("not meet any")
	ErrNotSupportFunc = errors.New("not support enumerable func")
)

func (m *Model) Exec(metrics []map[string]interface{}) (map[string]interface{}, error) {
	var res []bool
	for i := m.wr.low - 1; i <= m.wr.high-1; i++ {
		r, err := Evaluate(m.t, metrics[i])
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}

	andRes := res[0]
	orRes := res[0]

	for i := 1; i < len(res); i++ {
		andRes = andRes && res[i]
		orRes = orRes || res[i]
	}

	switch m.ef {
	case "any":
		fmt.Println("====>any", orRes, andRes)
		if orRes {
			return m.outputResult(metrics[0])
		}
		return nil, ErrNotMeetAny
	case "none":
		fmt.Println("====>none", orRes, andRes)
		if andRes == false {
			return m.outputResult(metrics[0])
		}
		return nil, ErrNotMeetNone
	case "each":
		fmt.Println("====>each", orRes, andRes)
		if andRes == true {
			return m.outputResult(metrics[0])
		}
		return nil, ErrNotMeetEach
	default:
		return nil, ErrNotSupportFunc
	}
}

// precedence：logicalOp > comparisonOp > arithmeticOp > unaryOp

func IsLogicalOp(v Value) bool {
	t := v.Type()

	if t != "binop" {
		return false
	}

	val, ok := v.Value().(string)
	if !ok {
		return false
	}

	if val == "and" || val == "or" {
		return true
	}

	return false
}

func IsUnaryOp(v Value) bool {
	return v.Type() == "unaryop"
}

func IsComparisonOp(v Value) bool {
	t := v.Type()
	if t != "binop" {
		return false
	}

	val, ok := v.Value().(string)
	if !ok {
		return false
	}

	switch val {
	case "<", "<=", ">", ">=", "==", "!=":
		return true
	default:
		return false
	}
}

type BinaryOperator struct { // "and", "or", "+"...
	Val string
}

func (o *BinaryOperator) Type() string {
	return "binop"
}

func (o *BinaryOperator) Value() interface{} {
	return o.Val
}

type UnaryOperator struct { // roundDown, roundUp, abs
	Val string
}

func (o *UnaryOperator) Type() string {
	return "unaryop"
}

func (o *UnaryOperator) Value() interface{} {
	return o.Val
}

type Literal struct { // 500, 3.14, "on"
	Val interface{}
}

func (l *Literal) Type() string {
	return "literal"
}

func (l *Literal) Value() interface{} {
	return l.Val
}

type Variable struct { // "speed"
	Val string
}

func (v *Variable) Type() string {
	return "variable"
}

func (v *Variable) Value() interface{} {
	return v.Val
}

// semantic tree
type Tree interface {
	GetParent() Tree
	SetParent(Tree)
	GetValue() Value
	SetLeftChild(Tree) Tree
	GetLeftChild() Tree
	SetRightChild(Tree) Tree
	GetRightChild() Tree
}

type Value interface {
	Type() string
	Value() interface{}
}

// Node is an implementation of Tree
// and each node can be seen as a tree
type Node struct {
	V Value
	l *Node // left node
	r *Node // right node
	p *Node // parent node
}

func (t *Node) GetParent() Tree {
	return t.p
}

func (t *Node) SetParent(pt Tree) {
	t.p = pt.(*Node)
}

func (t *Node) GetValue() Value {
	return t.V
}

func (t *Node) SetLeftChild(lt Tree) Tree {
	n := lt.(*Node)
	t.l = n
	n.p = t
	return lt
}

func (t *Node) GetLeftChild() Tree {
	return t.l
}

func (t *Node) SetRightChild(rt Tree) Tree {
	n := rt.(*Node)
	t.r = n
	n.p = t
	return rt
}

func (t *Node) GetRightChild() Tree {
	return t.r
}

// pre order traverse
func preOrderTraverse(t *Node, level int, enterF func(*Node, int), exitF func(*Node, int)) {
	if t == nil {
		return
	}

	if enterF != nil {
		enterF(t, level) // traverse this node
	}

	// traverse left children
	preOrderTraverse(t.l, level+1, enterF, exitF)

	// traverse right children
	preOrderTraverse(t.r, level+1, enterF, exitF)

	if exitF != nil {
		exitF(t, level) // traverse this node again
	}
}

// post order traverse
func postOrderTraverse(t *Node, level int, enterF func(*Node, int), exitF func(*Node, int)) {
	if t == nil {
		return
	}

	// traverse left children
	postOrderTraverse(t.l, level+1, enterF, exitF)

	// traverse right children
	postOrderTraverse(t.r, level+1, enterF, exitF)

	if enterF != nil {
		enterF(t, level) // traverse this node
	}
	if exitF != nil {
		exitF(t, level) // traverse this node again
	}
}

// in order traverse
func inOrderTraverse(t *Node, level int, enterF func(*Node, int), exitF func(*Node, int)) {
	if t == nil {
		return
	}

	// traverse left children
	inOrderTraverse(t.l, level+1, enterF, exitF)

	if enterF != nil {
		enterF(t, level) // traverse this node
	}

	// traverse right children
	inOrderTraverse(t.r, level+1, enterF, exitF)

	if exitF != nil {
		exitF(t, level) // traverse this node again
	}
}

func printPrefix(level int) {
	for i := 0; i < level; i++ {
		if i == level-1 {
			fmt.Printf(" |---")
		} else {
			fmt.Printf("     ")
		}
	}
}

func Dump(t Tree, order string) {
	var f = func(n *Node, level int) {
		if n == nil {
			return
		}

		printPrefix(level)

		if n.p == nil {
			// root node
			fmt.Printf("[root]()\n")
		} else {
			fmt.Printf("[%s](%v)\n", n.V.Type(), n.V.Value())
		}
	}

	switch order {
	default:
		// preorder
		preOrderTraverse(t.(*Node), 0, f, nil)
	case "inorder":
		inOrderTraverse(t.(*Node), 0, f, nil)
	case "postorder":
		postOrderTraverse(t.(*Node), 0, f, nil)
	}
}

func Evaluate(t Tree, m map[string]interface{}) (result bool, err error) {
	var s Stack[Value]

	defer func() {
		// extract error from panic
		if x := recover(); x != nil {
			result, err = false, fmt.Errorf("eval error: %v", x)
			return
		}
	}()

	var enterF = func(n *Node, level int) {
		if n == nil {
			return
		}

		if n.p == nil {
			// root node
			return
		}

		fmt.Printf("enter node: %v\n", n.GetValue().Value())
	}

	var exitF = func(n *Node, level int) {
		if n == nil {
			return
		}

		if n.p == nil {
			// root node
			return
		}

		fmt.Printf("exit node: %v\n", n.GetValue().Value())

		v := n.GetValue()
		switch v.Type() {
		case "binop":
			rhs, lhs := s.Pop(), s.Pop()
			s.Push(evalBinaryOpExpr(v.Value().(string), lhs, rhs))
		case "unaryop":
			lhs := s.Pop()
			s.Push(evalUnaryOpExpr(v.Value().(string), lhs))
		case "literal":
			s.Push(v)
		case "variable":
			name := v.Value().(string)
			value, ok := m[name]
			if !ok {
				panic(fmt.Sprintf("not found variable: %s", name))
			}

			// use the value in map to replace variable
			s.Push(&Literal{
				Val: value,
			})
		}
	}

	preOrderTraverse(t.(*Node), 0, enterF, exitF)
	result = s.Pop().Value().(bool)
	return
}

func lessThan(left, right interface{}) bool {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l < r
	case string:
		r := right.(string)
		return l < r
	case float64:
		r := right.(float64)
		return l < r
	}

	return false
}

func lessThanOrEqual(left, right interface{}) bool {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l <= r
	case string:
		r := right.(string)
		return l <= r
	case float64:
		r := right.(float64)
		return l <= r
	}
	return false
}

func largeThan(left, right interface{}) bool {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l > r
	case string:
		r := right.(string)
		return l > r
	case float64:
		r := right.(float64)
		return l > r
	}
	return false
}

func largeThanOrEqual(left, right interface{}) bool {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l >= r
	case string:
		r := right.(string)
		return l >= r
	case float64:
		r := right.(float64)
		return l >= r
	}
	return false
}

func equal(left, right interface{}) bool {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l == r
	case string:
		r := right.(string)
		return l == r
	case float64:
		r := right.(float64)
		return l == r
	case bool:
		r := right.(bool)
		return l == r
	}
	return false
}

func notEqual(left, right interface{}) bool {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l != r
	case string:
		r := right.(string)
		return l != r
	case float64:
		r := right.(float64)
		return l != r
	case bool:
		r := right.(bool)
		return l != r
	}
	return false
}

func add(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l + r
	case string:
		r := right.(string)
		return l + r
	case float64:
		r := right.(float64)
		return l + r
	default:
		panic("unsupport operand type in add")
	}
}

func sub(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l - r
	case float64:
		r := right.(float64)
		return l - r
	default:
		panic("unsupport operand type in sub")
	}
}

func multi(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l * r
	case float64:
		r := right.(float64)
		return l * r
	default:
		panic("unsupport operand type in multi")
	}
}

func divide(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l / r
	case float64:
		r := right.(float64)
		return l / r
	default:
		panic("unsupport operand type in divide")
	}
}

func mod(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int:
		r := right.(int)
		return l % r
	default:
		panic("unsupport operand type in mod")
	}
}

func and(left, right interface{}) interface{} {
	switch l := left.(type) {
	case bool:
		r := right.(bool)
		return l && r
	default:
		panic("unsupport operand type in and")
	}
}

func or(left, right interface{}) interface{} {
	switch l := left.(type) {
	case bool:
		r := right.(bool)
		return l || r
	default:
		panic("unsupport operand type in or")
	}
}

func evalUnaryOpExpr(unaryOp string, lhs Value) Value {
	switch unaryOp {
	case "roundUp":
		f, ok := lhs.Value().(float64)
		if !ok {
			panic(fmt.Sprintf("only support float64 unaryOp operand type"))
		}
		return &Literal{
			Val: math.Ceil(f),
		}
	case "roundDown":
		f, ok := lhs.Value().(float64)
		if !ok {
			panic(fmt.Sprintf("only support float64 unaryOp operand type"))
		}
		return &Literal{
			Val: math.Floor(f),
		}
	case "abs":
		i, ok := lhs.Value().(int)
		if !ok {
			panic(fmt.Sprintf("only support int unaryOp operand type"))
		}
		if i < 0 {
			i = -i
		}
		return &Literal{
			Val: i,
		}
	default:
		panic(fmt.Sprintf("unsupport unaryOp %s", unaryOp))
	}
}

func evalBinaryOpExpr(binOp string, lhs, rhs Value) Value {
	switch binOp {

	case "<":
		result := lessThan(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "<=":
		result := lessThanOrEqual(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case ">":
		result := largeThan(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case ">=":
		result := largeThanOrEqual(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "==":
		result := equal(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "!=":
		result := notEqual(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "+":
		result := add(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "-":
		result := sub(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "*":
		result := multi(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "/":
		result := divide(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "%":
		result := mod(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "and":
		result := and(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	case "or":
		result := or(lhs.Value(), rhs.Value())
		return &Literal{
			Val: result,
		}
	}

	return &Literal{
		Val: false,
	}
}

// the node is root node when node.p is nil
func New() Tree {
	return &Node{}
}

// construct a tree based on a reversePolishExpr
func NewFrom(reversePolishExpr []Value) Tree {
	var s Stack[Tree]
	for _, v := range reversePolishExpr {
		switch v.Type() {
		case "literal", "variable":
			s.Push(&Node{
				V: v,
			})
		case "binop":
			rchild, lchild := s.Pop(), s.Pop()
			n := &Node{
				V: v,
			}
			n.SetLeftChild(lchild)
			n.SetRightChild(rchild)
			s.Push(n)
		case "unaryop":
			lchild := s.Pop()
			n := &Node{
				V: v,
			}
			n.SetLeftChild(lchild)
			s.Push(n)
		}

	}
	first := s.Pop()
	root := &Node{}
	root.SetLeftChild(first)
	return root
}
