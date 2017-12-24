package main

import "fmt"

func main(){
  var plants = []float32{1,2,3,4,5,6}

  var grid float32 = 12
  capacity := calculateCapacity(plants)
  fmt.Println("The capacity is", capacity)
  fmt.Println("Load is", (capacity/grid))
}


func calculateCapacity(array []float32) (capacity float32){
  for i:=0; i < len(array); i++ {
    capacity += array[i]
  }
  return capacity
}
