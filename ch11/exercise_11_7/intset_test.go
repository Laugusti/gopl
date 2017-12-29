package intset

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Laugusti/gopl/ch6/exercise_6_5/intset"
)

var rng = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func BenchmarkHas(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	for i := 0; i < b.N; i++ {
		var set intset.IntSet
		set.Has(x)
	}
}

func BenchmarkAdd(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	for i := 0; i < b.N; i++ {
		var set intset.IntSet
		set.Add(x)
	}
}

func BenchmarkAddAll(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	for i := 0; i < b.N; i++ {
		var set intset.IntSet
		set.AddAll([]int{x}...)
	}
}

func BenchmarkUnionWith(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	var set1, set2 intset.IntSet
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intset.IntSet) {
			s.UnionWith(&t)
		}(set1, set2)
	}
}

func BenchmarkIntersectWith(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	var set1, set2 intset.IntSet
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intset.IntSet) {
			s.IntersectWith(&t)
		}(set1, set2)
	}
}

func BenchmarkDifferenceWith(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	var set1, set2 intset.IntSet
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intset.IntSet) {
			s.DifferenceWith(&t)
		}(set1, set2)
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	var set1, set2 intset.IntSet
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intset.IntSet) {
			s.SymmetricDifference(&t)
		}(set1, set2)
	}
}

func BenchmarkLen(b *testing.B) {
	var set intset.IntSet
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.Len()
	}
}

func BenchmarkRemove(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	var set intset.IntSet
	set.Add(x)
	for i := 0; i < b.N; i++ {
		func(s intset.IntSet) {
			s.Remove(x)
		}(set)
	}
}

func BenchmarkClear(b *testing.B) {
	var set intset.IntSet
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkCopy(b *testing.B) {
	var set intset.IntSet
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.Copy()
	}
}

func BenchmarkElems(b *testing.B) {
	var set intset.IntSet
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.Elems()
	}
}

func BenchmarkString(b *testing.B) {
	var set intset.IntSet
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.String()
	}
}
