// my first code in golang. is much more a implementation of fibonacci than a hello world per se.

package main
import "fmt"

func fibonacci(max_value int){
  var first_element int = 0
  var second_element int = 1

  for first_element < max_value {
    fmt.Println(first_element)
    first_element = second_element - first_element
    second_element += first_element
  }
}

func main() {
    fibonacci(22)
}
