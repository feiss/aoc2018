package main

import (
  "os"
  "bufio"
  "fmt"
  "sort"
  "strconv"
)

type Action int
const (
  BEGINS Action = iota
  SLEEPS
  WAKESUP
)


type Timestamp struct {
  day string
  hour int
  minute int
  guard string
  action Action
}

func readTimestamps(file string) ([]Timestamp) {
  f, _ := os.Open(file)
  defer f.Close()
  scanner := bufio.NewScanner(f)
  var timestamps []Timestamp
  for scanner.Scan() {
    var ts Timestamp
    var str1, str2 string
    fmt.Sscanf(scanner.Text(), "[%s %d:%d] %s %s", &ts.day, &ts.hour, &ts.minute, &str1, &str2)
    if str2[0:1] == "#" {
      ts.guard = str2[1:]
    } else {
      switch str2{
        case "asleep": ts.action = SLEEPS
        case "up": ts.action = WAKESUP
        default: ts.action = BEGINS
      }
    }
    timestamps = append(timestamps, ts)
  }
  return timestamps
}

func main() {
  var timestamps []Timestamp = readTimestamps("input.txt")

  sort.Slice(timestamps, func(i, j int) bool {
    if timestamps[i].day < timestamps[j].day { return true }
    if timestamps[i].day > timestamps[j].day { return false }
    if timestamps[i].hour > timestamps[i].hour { return true }
    if timestamps[i].hour < timestamps[i].hour { return false }
    return timestamps[i].minute < timestamps[j].minute
  })

  minutesAsleep := make(map[string][]int)
  var guard string
  var sleptAt int

  for _, v := range timestamps {
    switch v.action {
        case BEGINS: guard = v.guard
        case SLEEPS: sleptAt = v.minute
        case WAKESUP:
          for i:= sleptAt; i < v.minute; i++ {
            minutesAsleep[guard] = append(minutesAsleep[guard], i)
          }
      }
  }

  // find sleepest guard

  var maxSleep int
  var sleepestGuard string
  for k, v := range minutesAsleep {
    if len(v) > maxSleep {
      maxSleep = len(v)
      sleepestGuard = k
    }
  }

  // found sleepest guard, find the most frequent sleep minute
  var minutes []int = make([]int, 60)
  for _, v := range minutesAsleep[sleepestGuard] {
    minutes[v] ++
  }
  sort.Slice(minutes, func(i, j int) bool {
    return minutes[i] > minutes[j]
  })

  g, _ := strconv.ParseInt(sleepestGuard, 10, 32)
  fmt.Println(int64(minutes[0]) * g)
 }
