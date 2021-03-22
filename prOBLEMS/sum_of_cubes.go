package kata

func SumCubes(n int) int {
  // your code here
    sum:=0
    for i:=1; i<=n; i++ { 
        sum += (i*i*i)
    }
  return sum
}