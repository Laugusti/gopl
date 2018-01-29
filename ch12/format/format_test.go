package format

import (
	"fmt"
	"testing"
	"time"

	"gopl.io/ch12/format"
)

func TestAny(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))                  // "1"
	fmt.Println(format.Any(d))                  // "1"
	fmt.Println(format.Any([]int64{x}))         // "[]int64 0xc420010170"
	fmt.Println(format.Any([]time.Duration{d})) // "[]time.Duration 0xc420010178"
}
