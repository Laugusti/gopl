package format

import (
	"fmt"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	fmt.Println(Any(2 * time.Nanosecond))
	fmt.Println(Any(time.Now))
	fmt.Println(Any(true))
}
