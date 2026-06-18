package kata
​
import "strings"
​
func ReverseWords(str string) string {
  newStr := strings.Split(str, " ")
  
  for index, char := range newStr {
    result := ""
    
    for _, s := range char {
      result = string(s) + result
    }
    newStr[index] = result
  }
  return strings.Join(newStr, " ")
}