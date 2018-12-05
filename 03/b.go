package main

import (
  "os"
  "bufio"
  "fmt"
)

type Claim struct {
  id uint
  x, y, w, h uint
}

func readInput(file string) ([]Claim) {
  f, _ := os.Open(file)
  defer f.Close()
  scanner := bufio.NewScanner(f)
  var claims []Claim
  for scanner.Scan() {
    var claim Claim
    fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &claim.id, &claim.x, &claim.y, &claim.w, &claim.h)
    claims = append(claims, claim)
  }
  return claims
}

func drawClaim(fabric []uint, claim Claim) {
  for y := claim.y; y < claim.y + claim.h; y++ {
    for x := claim.x; x < claim.x + claim.w; x++ {
      fabric[y * 1000 + x] ++
    }
  }
}

func checkClaim(fabric []uint, claim Claim) (bool) {
  for y := claim.y; y < claim.y + claim.h; y++ {
    for x := claim.x; x < claim.x + claim.w; x++ {
      if fabric[y * 1000 + x] > 1 {
        return false
      }
    }
  }
  return true
}

func main() {
  var claims []Claim = readInput("input.txt")
  var fabric []uint = make([]uint, 1000000, 1000000)

  for _,c := range claims {
    drawClaim(fabric, c)
  }

  for _,c := range claims {
    if checkClaim(fabric, c) {
      fmt.Println("Claim #", c.id)
      return
    }
  }

  fmt.Println("Claim not found")
}

