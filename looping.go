package main

import "fmt"

func main() {
  slice := []int{1, 2, 3}

  for i:=0; i < 5; i++ {
    fmt.Println(slice[i])
    slice = append(slice, i)
    fmt.Println(slice)
  }

  myMap := make(map[int]string)
  myMap[0]="Samsung"
  myMap[1]="Apple"

for i, v := range myMap {
  fmt.Println(i,v)
}

}
