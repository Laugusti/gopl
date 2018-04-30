package display

import (
	"testing"
)

type cyclicStruct struct {
	cs interface{}
}

func TestCyclicStruct(t *testing.T) {
	cs := cyclicStruct{}
	cs.cs = &cs
	Display("cs", cs)
}
