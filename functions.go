package main

import (
  "fmt"
)

func main(){
  returned, terms := sum(1,2,3)
  fmt.Println(returned, terms)
}

func sum(terms ...int) (int, result int){
  //result := 0 // Initialized when calling the variable in the returned value
  for _, term:= range terms{
    result += term
  }
  return result, len(terms)
}
