package kata

func Solution(word string) string {
  var result = ""
  for _,c := range word {
    result = string(c) + result
  }
  return result
}