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
  codes := readFileToArray("input.txt")

  for ia := 0; ia < len(codes) - 1; ia++ {
    for ib := ia + 1; ib < len(codes); ib++ {

      different := -1
      for i := 0; i < len(codes[ia]); i++ {
        if codes[ia][i] != codes[ib][i] {
          if different == -1 {
            different = i
          } else {
            goto notFound
          }
        }
      }
      // found
      fmt.Println(codes[ia][:different] + codes[ia][different + 1:])
      return

      notFound:
    }
  }
}
