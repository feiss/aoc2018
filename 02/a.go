package main

import (
  "os"
  "bufio"
  "fmt"
)

func readFileToArray(file string) ([]string) {
  f, _ := os.Open(file)
  defer f.Close()
  scanner := bufio.NewScanner(f)
  var arr []string
  for scanner.Scan() {
    arr = append(arr, scanner.Text())
  }
  return arr
}


func main() {
  arr := readFileToArray("input.txt")
  var counter [128]byte
  twos := 0
  threes := 0
  for _, code := range arr {

    hasTwos := false
    hasThrees := false

    for i := 0; i < len(code); i++ { counter[code[i]] ++ }

    for i := 0; i < len(counter); i++ {
      if counter[i] == 2 { hasTwos = true}
      if counter[i] == 3 { hasThrees = true}
      counter[i] = 0
    }

    if hasTwos { twos ++ }
    if hasThrees { threes ++ }
  }
  fmt.Println("Checksum:", twos * threes)
}
