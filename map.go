package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    f := strings.Fields(s)
    for _, str := range f {
    	m[str]++
    }
    
    return m
}

func main() {
    wc.Test(WordCount)
}
