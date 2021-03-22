package kata

func GetCount(str string) (count int) {
  for _, n := range str {
    switch n{
      case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}