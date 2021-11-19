package main
import "fmt"
 
func coolFunction() {
  var coolVariable = 5
  fmt.Println(coolVariable)
}
 
func concat_strings(str1 string, str2 string) string {
    if str1 == "" {
        return str2 
    } else if str2 == "" {
        return str1 
    } else {
        return str1 + str2
    }
}
func main() {
  coolFunction()
  var x, y, z string 
  y = "ab"
  z = "cd"
  x = concat_strings(y, z)
  fmt.Println(x)
}
