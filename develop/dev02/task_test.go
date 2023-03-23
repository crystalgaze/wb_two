package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	m := make(map[string]string)
	m["q4w5e6r6t7y"] = "qqqqwwwwweeeeeerrrrrrttttttty"
	m["qwerty"] = "qwerty"
	m["42"] = ""
	m[""] = ""
	m["weuiito\\\\5\\3\\hwg\\ji5hfsdgwer7i"] = "weuiito\\\\\\\\\\3hwgjiiiiihfsdgwerrrrrrri"

	for k, v := range m {
		if s, _ := unpack(k); s != v {
			t.Error("Error in string: ", k)
		}
	}
}
