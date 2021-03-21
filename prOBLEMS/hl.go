
package kata

import (
  "math"
  "strconv"
  "strings"
)

func HighAndLow(i string) string {
  max := math.MinInt32
  min := math.MaxInt32
  for _, value := range strings.Split(i, " ") {
    num, _ := strconv.Atoi(value)
    if num > max {
      max = num
    }
    if num < min {
      min = num
    }
  }

  return strconv.Itoa(max) + " " + strconv.Itoa(min)
}
