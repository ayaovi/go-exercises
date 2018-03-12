package main

import (
  "fmt"
	"strings"
	"unicode"
  //"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
  words := strings.FieldsFunc(s, func(c rune) bool {
    return !unicode.IsLetter(c) && !unicode.IsNumber(c)
  })
  
  mappings := make(map[string]int)
  
  for _, value := range words {
    _, ok := mappings[value]
    if ok {
      mappings [value] += 1
    } else {
      mappings [value] = 1
    }
  }
  
  return mappings
}

func main() {
  //wc.Test(WordCount)
  fmt.Println(WordCount("Once upon a time, a long long time again, I was there and forever will be."))
}
