package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressedString(word string) string {
	var comp []string
	var compReady string
	var count int
	buf := rune(word[0])
	for _, i := range word {
		if buf != i {
			comp = append(comp, strconv.Itoa(count))
			comp = append(comp, string(buf))
			buf = i
			count = 1
			continue
		}
		if count == 9 {
			comp = append(comp, strconv.Itoa(count))
			comp = append(comp, string(buf))
			count = 1
		} else {
			count++
		}
	}
	comp = append(comp, strconv.Itoa(count))
	comp = append(comp, string(buf))
	compReady = strings.Join(comp, "")
	return compReady
}

func main() {
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))
}
