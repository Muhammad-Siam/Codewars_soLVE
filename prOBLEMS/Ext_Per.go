package kata

func ExtraPerfect(n int) []int {
  var r []int
  for i := 1; i <= n; i += 2 {
    r = append(r, i)
  }
  return r
}
