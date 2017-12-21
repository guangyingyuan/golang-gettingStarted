package main

import "fmt"

func main(){
  var plants = []float32{1,2,3,4,5,6}
  var capacity float32 = 0

  for i:=0; i < len(plants); i++ {
    capacity += plants[i]
  }
  var grid float32 = 12
  fmt.Println("The capacity is", capacity)
  fmt.Println("Load is", (capacity/grid))
}
