package main

import (
  "fmt"
	"strings"
  //"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
  words := strings.Fields(s)
  
  mappings := map[string]int{}
  
  for _, value := range words {
      mappings [value]++
  }
  
  return mappings
}

func main() {
  //wc.Test(WordCount)
  fmt.Println(WordCount("Once upon a time, a long long time again, I was there and forever will be."))
}
