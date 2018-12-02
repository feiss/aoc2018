package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func valueInSlice(list []int, value int) (bool) {
  for _, v := range list {
    if value == v { return true }
  }
  return false
}

func main() {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  var val int = 0
  var changes []int
  for scanner.Scan() {
    i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
    changes = append(changes, int(i))
  }

  var frecs []int
  var i int = 0
  for {
    val += changes[i]
    if valueInSlice(frecs, val) {
      fmt.Println(val)
      return
    }
    frecs = append(frecs, val)
    i = (i + 1) % len(changes)
  }
}
