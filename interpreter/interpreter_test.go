package interpreter

import "testing"

func TestInterpreter (t *testing.T) {
	context.judge("Peter is a unmarried man")
	context.judge("May is a married woman")
}
