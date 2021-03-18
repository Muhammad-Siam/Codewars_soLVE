package kata

import "strings"

func Accum(s string) string {
    parts := make([]string, len(s))
    
    for i := 0; i < len(s); i++ {
      parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
    }
    
    return strings.Join(parts, "-")
}