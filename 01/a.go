package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  val := 0
  for scanner.Scan() {
    i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
    val += int(i)
  }
  fmt.Println(val)
}
