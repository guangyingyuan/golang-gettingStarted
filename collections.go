package main

import "fmt"

const (
  first = 1 << (10 * iota)
  second
  third
)
const (
  fourth = iota // it's zero!
)

func main(){
  println(first,"\n",second,"\n", third,"\n",fourth)

  myArray := [...]string{"Blue", "Yelow", "Red"}
  fmt.Println(myArray)

  //Copy an Array (creating a slice)
  /*
  Arrays have their place, but they're a bit inflexible, so you don't see them too often in Go code.
  Slices, though, are everywhere. They build on arrays to provide great power and convenience.
  */
  mySlice := myArray[:]
  fmt.Println(mySlice)

  //Increment Array
  mySlice = append(mySlice, "Brown")
  mySlice = append(mySlice, "Pink")
  fmt.Println(mySlice, len(mySlice))

  //Creating a Slice

  newSlice := []float32{1., 2.1, second}
  fmt.Println(newSlice)
  newSlice = append(newSlice, 2.3)
  fmt.Println(newSlice)

  // Making the slice
  lastSlice := make([]float32, 10); // This add zero to the empty possitions
  lastSlice[0] = 12.
  lastSlice[9] = 21.
  fmt.Println(lastSlice)


  //Mapping

  myMap := make(map[int]string)
  myMap[1] = "Mercedes"
  myMap[14] = "Ford"
  fmt.Println(myMap)
  fmt.Println(myMap[1])
  fmt.Println(myMap[0])
}
