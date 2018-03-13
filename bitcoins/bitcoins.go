// You have 50 bitcoins to distribute to 10 users: 
// Matthew, Sarah, Augustus, Heidi, Emilie, Peter, 
// Giana, Adriano, Aaron, Elizabeth The coins will 
// be distributed based on the vowels contained in 
// each name where:
//
// a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins
//
// and a user can’t get more than 10 coins. Print 
// a map with each user’s name and the amount of 
// coins distributed. After distributing all the 
// coins, you should have 2 coins left.

package main

import (
  "fmt"
  "strings"
)

func CalculateBTC(name string) int {
  btc := 0
  for _, character := range strings.ToLower(name) {
    switch character {
      case 'a':
        btc++
      case 'e':
        btc++
      case 'i':
        btc += 2
      case 'o':
        btc += 3
      case 'u':
        btc += 4
    }
  }
  if btc > 10 { return 10 }
  return btc
}

func main() {
  totalBtc := 50
  names := []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
  btcForUsers := map[string]int{}
  
  for _, name := range names {
    btc := CalculateBTC(name)
    btcForUsers[name] = btc
    totalBtc -= btc
  }
  
  fmt.Println(btcForUsers)
  fmt.Println("Coins left: ", totalBtc)
}