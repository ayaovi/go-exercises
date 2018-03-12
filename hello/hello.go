package main

import "fmt"

func main() {
  //pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
  names := []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}
  
  max_length := 0
  for _, v:= range names {
    if l := len(v); l > max_length {
      max_length = l
    }
  }
  
  res := make([][]string, max_length)
  
  for _, v:= range names {
    res[len(v) - 1] = append(res[len(v) - 1], v)
  }
  fmt.Println(res)
}