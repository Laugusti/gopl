package intset

import (
	"testing"

	"github.com/Laugusti/gopl/ch11/exercise_11_7/intmap"
)

func BenchmarkMapHas(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	for i := 0; i < b.N; i++ {
		set := make(intmap.IntMap)
		set.Has(x)
	}
}

func BenchmarkMapAdd(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	for i := 0; i < b.N; i++ {
		set := make(intmap.IntMap)
		set.Add(x)
	}
}

func BenchmarkMapAddAll(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	for i := 0; i < b.N; i++ {
		set := make(intmap.IntMap)
		set.AddAll([]int{x}...)
	}
}

func BenchmarkMapUnionWith(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	set1, set2 := make(intmap.IntMap), make(intmap.IntMap)
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intmap.IntMap) {
			s.UnionWith(&t)
		}(set1, set2)
	}
}

func BenchmarkMapIntersectWith(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	set1, set2 := make(intmap.IntMap), make(intmap.IntMap)
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intmap.IntMap) {
			s.IntersectWith(&t)
		}(set1, set2)
	}
}

func BenchmarkMapDifferenceWith(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	set1, set2 := make(intmap.IntMap), make(intmap.IntMap)
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intmap.IntMap) {
			s.DifferenceWith(&t)
		}(set1, set2)
	}
}

func BenchmarkMapSymmetricDifference(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	set1, set2 := make(intmap.IntMap), make(intmap.IntMap)
	set1.Add(x)
	set2.Add(x)
	for i := 0; i < b.N; i++ {
		func(s, t intmap.IntMap) {
			s.SymmetricDifference(&t)
		}(set1, set2)
	}
}

func BenchmarkMapLen(b *testing.B) {
	set := make(intmap.IntMap)
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.Len()
	}
}

func BenchmarkMapRemove(b *testing.B) {
	x := rng.Intn(1e8) + 99999999
	set := make(intmap.IntMap)
	set.Add(x)
	for i := 0; i < b.N; i++ {
		func(s intmap.IntMap) {
			s.Remove(x)
		}(set)
	}
}

func BenchmarkMapClear(b *testing.B) {
	set := make(intmap.IntMap)
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkMapCopy(b *testing.B) {
	set := make(intmap.IntMap)
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.Copy()
	}
}

func BenchmarkMapElems(b *testing.B) {
	set := make(intmap.IntMap)
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.Elems()
	}
}

func BenchmarkMapString(b *testing.B) {
	set := make(intmap.IntMap)
	set.Add(rng.Intn(1e8) + 99999999)
	for i := 0; i < b.N; i++ {
		set.String()
	}
}
