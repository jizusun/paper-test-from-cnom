// Application which greets you.
package main

import "fmt"

func main() {
	result := noName("kubernetes", "ukbernetes")
	fmt.Println(result)
}

// TODO: benchmark
func noName(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	countA := countChar(a)
	countB := countChar(b)
	if len(countA) != len(countB) {
		return false
	}
	for key, elem := range countA {
		if elem != countB[key] {
			return false
		}
	}
	return true
}

func countChar(str string) map[string]int {
	count := make(map[string]int)
	for _, r := range str {
		ch := string(r)
		if val, ok := count[ch]; ok {
			count[ch] = val + 1
		} else {
			count[ch] = 1
		}
	}
	return count
}
