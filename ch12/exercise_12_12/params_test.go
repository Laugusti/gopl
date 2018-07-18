package params

import "testing"

func TestPack(t *testing.T) {
	tests := []struct {
		rawURL string
		input  interface{}
		want   string
	}{
		{"", struct{ a int }{3}, "?a=3"},
		{"", struct{ a []int }{[]int{1, 2, 3}}, "?a=1&a=2&a=3"},
		{"http://example.com", struct {
			a string `http:"b"`
		}{"testing"}, "http://example.com?b=testing"},
		{"https://example.com?l=English#fragment", struct {
			x        bool
			language string `http:"l"`
			maxSize  int    `http:"max"`
		}{language: "官话", maxSize: 100}, "https://example.com?l=English&l=%E5%AE%98%E8%AF%9D&max=100&x=false#fragment"},
	}

	for _, test := range tests {
		got, err := Pack(test.rawURL, test.input)
		if err != nil {
			t.Fatal(err)
		}
		if got != test.want {
			t.Errorf("TestPack => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}
