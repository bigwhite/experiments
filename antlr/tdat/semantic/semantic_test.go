package semantic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tree := New()

	// and
	n1 := tree.SetLeftChild(&Node{
		V: &BinaryOperator{
			Val: "or",
		},
	}).SetLeftChild(&Node{
		V: &BinaryOperator{
			Val: "and",
		},
	})

	n1.SetLeftChild(&Node{
		V: &BinaryOperator{
			Val: "<",
		},
	})

	n1.GetLeftChild().SetLeftChild(&Node{
		V: &Variable{
			Val: "speed",
		},
	})

	n1.GetLeftChild().SetRightChild(&Node{
		V: &Literal{
			Val: 50,
		},
	})

	n2 := n1.SetRightChild(&Node{
		V: &BinaryOperator{
			Val: "<",
		},
	})

	n2.SetLeftChild(&Node{
		V: &Variable{
			Val: "temperature",
		},
	})

	n2.SetRightChild(&Node{
		V: &Literal{
			Val: 2,
		},
	})

	n3 := tree.GetLeftChild().SetRightChild(&Node{
		V: &BinaryOperator{
			Val: "<",
		},
	})

	n3.SetLeftChild(&Node{
		V: &UnaryOperator{
			Val: "roundDown",
		},
	}).SetLeftChild(&Node{
		V: &Variable{
			Val: "salinity",
		},
	})

	n3.SetRightChild(&Node{
		V: &Literal{
			Val: 600.0,
		},
	})

	Dump(tree, "preorder")

	fmt.Println("")

	Dump(tree, "postorder")
	//($speed < 50) and ($temperature < 2) or (roundDown($salinity) < 600.0);
	m := map[string]interface{}{
		"speed":       30,
		"salinity":    500.0,
		"temperature": 1,
	}
	r, err := Evaluate(tree, m)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}
	fmt.Printf("result = %v\n", r)

}

func newBinaryOperator(v string) Value {
	return &BinaryOperator{
		Val: v,
	}
}

func newUnaryOperator(v string) Value {
	return &UnaryOperator{
		Val: v,
	}
}

func newVariable(v string) Value {
	return &Variable{
		Val: v,
	}
}

func newLiteral(v interface{}) Value {
	return &Literal{
		Val: v,
	}
}

func TestNewFrom(t *testing.T) {
	//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
	// speed,50,<,temperature,1,+,4,<,and,salinity,roundDown,600,<=,ph,roundUp,8.0,>,or,or
	var reversePolishExpr []Value

	reversePolishExpr = append(reversePolishExpr, newVariable("speed"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(50))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newVariable("temperature"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(1))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("+"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(4))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("and"))
	reversePolishExpr = append(reversePolishExpr, newVariable("salinity"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundDown"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(600.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<="))
	reversePolishExpr = append(reversePolishExpr, newVariable("ph"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundUp"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(8.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator(">"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))

	tree := NewFrom(reversePolishExpr)
	Dump(tree, "preorder")

	// test table
	var cases = []struct {
		id       string
		m        map[string]interface{}
		expected bool
	}{
		//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
		{
			id: "0001",
			m: map[string]interface{}{
				"speed":       30,
				"temperature": 6,
				"salinity":    700.0,
				"ph":          7.0,
			},
			expected: false,
		},
		{
			id: "0002",
			m: map[string]interface{}{
				"speed":       30,
				"temperature": 1,
				"salinity":    500.0,
				"ph":          7.0,
			},
			expected: true,
		},
		{
			id: "0003",
			m: map[string]interface{}{
				"speed":       60,
				"temperature": 10,
				"salinity":    700.0,
				"ph":          9.0,
			},
			expected: true,
		},
		{
			id: "0004",
			m: map[string]interface{}{
				"speed":       30,
				"temperature": 1,
				"salinity":    700.0,
				"ph":          9.0,
			},
			expected: true,
		},
	}

	for _, caze := range cases {
		r, err := Evaluate(tree, caze.m)
		if err != nil {
			t.Errorf("[case %s]: want nil, actual %s", caze.id, err.Error())
		}
		if r != caze.expected {
			t.Errorf("[case %s]: want %v, actual %v", caze.id, caze.expected, r)
		}
	}
}

func TestModelExecEach(t *testing.T) {
	//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
	// speed,50,<,temperature,1,+,4,<,and,salinity,roundDown,600,<=,ph,roundUp,8.0,>,or,or
	var reversePolishExpr []Value

	reversePolishExpr = append(reversePolishExpr, newVariable("speed"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(50))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newVariable("temperature"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(1))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("+"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(4))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("and"))
	reversePolishExpr = append(reversePolishExpr, newVariable("salinity"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundDown"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(600.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<="))
	reversePolishExpr = append(reversePolishExpr, newVariable("ph"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundUp"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(8.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator(">"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))

	model := NewModel(reversePolishExpr, WindowsRange{1, 3}, "each", []string{})

	// test table
	var cases = []struct {
		id       string
		m        []map[string]interface{}
		expected []string
	}{
		//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
		{
			id: "0001",
			m: []map[string]interface{}{
				{
					"speed":       30,
					"temperature": 6,
					"salinity":    500.0,
					"ph":          7.0,
				},
				{
					"speed":       31,
					"temperature": 7,
					"salinity":    501.0,
					"ph":          7.1,
				},
				{
					"speed":       30,
					"temperature": 6,
					"salinity":    498.0,
					"ph":          6.9,
				},
			},
			expected: []string{"speed", "temperature", "salinity", "ph"},
		},
	}

	for _, caze := range cases {
		r, err := model.Exec(caze.m)
		if err != nil {
			t.Errorf("[case %s]: want nil, actual %s", caze.id, err.Error())
		}

		if len(r) != len(caze.expected) {
			t.Errorf("[case %s]: want %v, actual %v", caze.id, len(caze.expected), len(r))
		}

		expected := caze.m[0]
		for _, item := range caze.expected {
			i, ok := r[item]
			if !ok {
				t.Errorf("[case %s]: want %s, actual not exist", caze.id, item)
			}

			if !reflect.DeepEqual(expected[item], i) {
				t.Errorf("[case %s]: want %v, actual %v", caze.id, expected[item], i)
			}
		}
	}
}

func TestModelExecNone(t *testing.T) {
	//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
	// speed,50,<,temperature,1,+,4,<,and,salinity,roundDown,600,<=,ph,roundUp,8.0,>,or,or
	var reversePolishExpr []Value

	reversePolishExpr = append(reversePolishExpr, newVariable("speed"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(50))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newVariable("temperature"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(1))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("+"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(4))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("and"))
	reversePolishExpr = append(reversePolishExpr, newVariable("salinity"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundDown"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(600.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<="))
	reversePolishExpr = append(reversePolishExpr, newVariable("ph"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundUp"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(8.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator(">"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))

	model := NewModel(reversePolishExpr, WindowsRange{1, 3}, "none", []string{})

	// test table
	var cases = []struct {
		id       string
		m        []map[string]interface{}
		expected []string
	}{
		//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
		{
			id: "0001",
			m: []map[string]interface{}{
				{
					"speed":       30,
					"temperature": 6,
					"salinity":    900.0,
					"ph":          7.0,
				},
				{
					"speed":       31,
					"temperature": 7,
					"salinity":    901.0,
					"ph":          7.1,
				},
				{
					"speed":       30,
					"temperature": 6,
					"salinity":    898.0,
					"ph":          7.2,
				},
			},
			expected: []string{"speed", "temperature", "salinity", "ph"},
		},
	}

	for _, caze := range cases {
		r, err := model.Exec(caze.m)
		if err != nil {
			t.Errorf("[case %s]: want nil, actual %s", caze.id, err.Error())
		}

		if len(r) != len(caze.expected) {
			t.Errorf("[case %s]: want %v, actual %v", caze.id, len(caze.expected), len(r))
		}

		expected := caze.m[0]
		for _, item := range caze.expected {
			i, ok := r[item]
			if !ok {
				t.Errorf("[case %s]: want %s, actual not exist", caze.id, item)
			}

			if !reflect.DeepEqual(expected[item], i) {
				t.Errorf("[case %s]: want %v, actual %v", caze.id, expected[item], i)
			}
		}
	}
}

func TestModelExecAny(t *testing.T) {
	//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
	// speed,50,<,temperature,1,+,4,<,and,salinity,roundDown,600,<=,ph,roundUp,8.0,>,or,or
	var reversePolishExpr []Value

	reversePolishExpr = append(reversePolishExpr, newVariable("speed"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(50))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newVariable("temperature"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(1))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("+"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(4))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("and"))
	reversePolishExpr = append(reversePolishExpr, newVariable("salinity"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundDown"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(600.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("<="))
	reversePolishExpr = append(reversePolishExpr, newVariable("ph"))
	reversePolishExpr = append(reversePolishExpr, newUnaryOperator("roundUp"))
	reversePolishExpr = append(reversePolishExpr, newLiteral(8.0))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator(">"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))
	reversePolishExpr = append(reversePolishExpr, newBinaryOperator("or"))

	model := NewModel(reversePolishExpr, WindowsRange{1, 3}, "any", []string{})

	// test table
	var cases = []struct {
		id       string
		m        []map[string]interface{}
		expected []string
	}{
		//($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0))
		{
			id: "0001",
			m: []map[string]interface{}{
				{
					"speed":       30,
					"temperature": 6,
					"salinity":    900.0,
					"ph":          7.0,
				},
				{
					"speed":       31,
					"temperature": 2,
					"salinity":    901.0,
					"ph":          7.1,
				},
				{
					"speed":       30,
					"temperature": 6,
					"salinity":    898.0,
					"ph":          7.2,
				},
			},
			expected: []string{"speed", "temperature", "salinity", "ph"},
		},
	}

	for _, caze := range cases {
		r, err := model.Exec(caze.m)
		if err != nil {
			t.Errorf("[case %s]: want nil, actual %s", caze.id, err.Error())
		}

		if len(r) != len(caze.expected) {
			t.Errorf("[case %s]: want %v, actual %v", caze.id, len(caze.expected), len(r))
		}

		expected := caze.m[0]
		for _, item := range caze.expected {
			i, ok := r[item]
			if !ok {
				t.Errorf("[case %s]: want %s, actual not exist", caze.id, item)
			}

			if !reflect.DeepEqual(expected[item], i) {
				t.Errorf("[case %s]: want %v, actual %v", caze.id, expected[item], i)
			}
		}
	}
}
