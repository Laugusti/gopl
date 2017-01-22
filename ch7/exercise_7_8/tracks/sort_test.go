package tracks

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func lessYear(t1, t2 *Track) bool {
	return t1.Year < t2.Year
}

func lessLength(t1, t2 *Track) bool {
	return t1.Length < t2.Length
}

func createTracks(n, m int) []*Track {
	t := make([]*Track, n)
	for i := 0; i < n; i++ {
		t[i] = &Track{"title", "artist", "album", rand.Intn(m), time.Duration(rand.Intn(m))}
	}
	return t
}

func inOrderByYearAndLength(t []*Track) bool {
	for i := 1; i < len(t); i++ {
		if t[i-1].Year != t[i].Year {
			if t[i-1].Year > t[i].Year {
				return false
			}
		} else if t[i-1].Length != t[i].Length {
			if t[i-1].Length > t[i].Length {
				return false
			}
		}
	}
	return true
}

func TestStableSort(t *testing.T) {
	n, m := 100000, 1000
	t1 := createTracks(n, m)

	if sort.IsSorted(customSort{t1, lessYear}) || sort.IsSorted(customSort{t1, lessLength}) {
		t.Errorf("terrible rand")
	}

	sort.Stable(customSort{t1, lessLength})
	sort.Stable(customSort{t1, lessYear})

	if !inOrderByYearAndLength(t1) {
		t.Errorf("TrackSorter is not multi-tiered")
	}
}
func TestTrackSorter(t *testing.T) {
	n, m := 100000, 1000
	t1 := createTracks(n, m)

	if sort.IsSorted(customSort{t1, lessYear}) || sort.IsSorted(customSort{t1, lessLength}) {
		t.Errorf("terrible rand")
	}

	sort.Sort(NewTrackSorter(t1, SortByLength))
	sort.Sort(NewTrackSorter(t1, SortByYear))

	if !inOrderByYearAndLength(t1) {
		t.Errorf("TrackSorter is not multi-tiered")
	}
}

func BenchmarkStableSort1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		t1 := createTracks(1000, 100)
		b.StartTimer()
		sort.Stable(customSort{t1, lessYear})
		b.StopTimer()
	}
}

func BenchmarkTrackSorter(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		t1 := createTracks(1000, 100)
		b.StartTimer()
		sort.Sort(NewTrackSorter(t1, SortByYear))
		b.StopTimer()
	}
}
