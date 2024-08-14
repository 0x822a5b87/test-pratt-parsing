package parser

import (
	"testing"
)

func TestParseExpression(t *testing.T) {
	expectedValues := []struct {
		input  string
		output string
	}{
		{`1`, `1`},
		{`-1`, `-1`},
		{`1 + 2 * 3`, `(1 + (2 * 3))`},
		{`2!`, `(2!)`},
		{`0 + 1 + 2! * -3`, `((0 + 1) + ((2!) * -3))`},
	}

	for i, expectedValue := range expectedValues {
		p := NewParser(expectedValue.input)
		expr := p.ParseExpression(0)
		p.Close()
		if expectedValue.output != expr.ToString() {
			t.Fatalf("offset = [%d], expected [%s], got [%s]", i, expectedValue.output, expr.ToString())
		}
	}
}
