package main

import (
 "fmt"
 "os"
 "bufio"
 "strconv"
 "strings"
)

func main() {
 scanner := bufio.NewScanner(os.Stdin)

 for {
  fmt.Print("enter number: ")
  scanner.Scan()
  input := scanner.Text()

  if input == "exit" {
   fmt.Println("end")
   break
  }

  expr := strings.Split(input, "")

  num1, err := strconv.ParseFloat(expr[0], 64)
  if err != nil {
   fmt.Println(err)
   continue
 }

  num2, err := strconv.ParseFloat(expr[2], 64)
  if err != nil {
   fmt.Println(err)
   continue
  }

  var result float64

  switch expr[1] {
  case "+":
   result = num1 + num2
  case "-":
   result = num1 - num2
  case "*":
   result = num1 * num2
  case "/":
   if num2 == 0 {
    fmt.Println("error: /0")
    continue
   }
   result = num1 / num2
  default:
   fmt.Println("error! enter: +, -, *, /")
   continue
  }

  fmt.Printf("%v\n", result)
 }
}
