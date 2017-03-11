package main

type linkdepth struct {
	link  string
	depth int
}

func makeLDSlice(links []string, depth int) []linkdepth {
	list := make([]linkdepth, len(links))
	for i := range links {
		list[i] = linkdepth{links[i], depth}
	}
	return list
}
