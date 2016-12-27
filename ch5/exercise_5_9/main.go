package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expand(string(bytes), strings.ToUpper))
}

func expand(s string, f func(string) string) string {
	for _, token := range getTokens(s) {
		s = strings.Replace(s, "$"+token, f(token), -1)
	}
	return s
}

func getTokens(s string) []string {
	tokens := make(map[string]bool)
	index := strings.Index(s, "$")
	if index < 0 {
		return nil
	}
	for _, s1 := range strings.Split(s[index+1:], "$") {
		//TODO: need to handle special chars
		t := getAlphaNumericWord(s1)
		if t != "" {
			tokens[t] = true
		}
	}
	return getKeysFromMap(tokens)
}

func getKeysFromMap(m map[string]bool) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return
}

func getAlphaNumericWord(s string) string {
	i := 0
	runes := []rune(s)
	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			i++
		}
	}
	return string(runes[:i])
}
