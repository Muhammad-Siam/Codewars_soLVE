package kata

func BaumSweet(ch chan<- int) {
  for i := 0; ; i++ {
    ch <- baum(i)
  }
}

func baum(n int) int {
  if n == 0 {
    return 1
  }
  for n%4 == 0 {
    n /= 4
  }
  if n&1 == 0 {
    return 0
  }
  return baum((n - 1) / 2)
}