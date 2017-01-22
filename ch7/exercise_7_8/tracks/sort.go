package tracks

type sortType int

// Represents Sort types as Constants
const (
	SortByTitle sortType = iota
	SortByArtist
	SortByAlbum
	SortByYear
	SortByLength
)

type trackSorter struct {
	Tracks       []*Track
	SortBy       sortType
	lastSortBy   sortType
	preSortOrder map[*Track]int
}

// NewTrackSorter returns a sort.Interface implementation for a slice of *Track by sortType
func NewTrackSorter(t []*Track, s sortType) *trackSorter {
	return &trackSorter{Tracks: t, SortBy: s}
}

// Implements the sort.Interface
func (s *trackSorter) Len() int      { return len(s.Tracks) }
func (s *trackSorter) Swap(i, j int) { s.Tracks[i], s.Tracks[j] = s.Tracks[j], s.Tracks[i] }
func (s *trackSorter) Less(i, j int) bool {
	// rebuild sort order for new sorts
	if s.lastSortBy != s.SortBy && len(s.preSortOrder) != len(s.Tracks) {
		s.buildPreSortOrder()
	}

	return less(s.Tracks[i], s.Tracks[j], s.preSortOrder, s.SortBy)
}

// buildPreSortOrder saves the order of the Track elements in a map
func (s *trackSorter) buildPreSortOrder() {
	s.preSortOrder = make(map[*Track]int)
	for i, t := range s.Tracks {
		s.preSortOrder[t] = i + 1
	}
	s.lastSortBy = s.SortBy
}

// returns true if x should precede y
func less(x, y *Track, preSortOrder map[*Track]int, s sortType) bool {
	switch s {
	case SortByTitle:
		if x.Title != y.Title {
			return x.Title < y.Title
		}
	case SortByArtist:
		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}
	case SortByAlbum:
		if x.Album != y.Album {
			return x.Album < y.Album
		}
	case SortByYear:
		if x.Year != y.Year {
			return x.Year < y.Year
		}
	case SortByLength:
		if x.Length != y.Length {
			return x.Length < y.Length
		}
	}
	return preSortOrder[x] < preSortOrder[y]
}
