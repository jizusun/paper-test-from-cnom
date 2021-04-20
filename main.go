// Application which greets you.
package main

func noName(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for x := range b {
		if a[0] == b[x] {
			return noName(utilityFunction(a, 0), utilityFunction(b, x))
		}
	}
	return len(b) == 0
}

func utilityFunction(s string, j int) string {
	ret := make([]rune, len(s)-1)
	d := 0
	for k := range s {
		if k == j {
			d = 1
		} else {
			ret[k-d] = rune(s[k])
		}
	}
	return string(ret)
}
