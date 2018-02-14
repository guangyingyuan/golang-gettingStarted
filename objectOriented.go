package main

import (
  "fmt"
)

type car struct{
  model string
  color string
  year string
}


func main(){

  myCar := car{}
  myCar.model = "Ford Focus"
  myCar.color = "Red"
  myCar.year = "2002"

  fmt.Println("My car is a", myCar.model, myCar.year)
  fmt.Println(myCar)

  mySecondCar := car {"Fiat 500", "White", "2017"}
  fmt.Println(mySecondCar)


  //With pointers
  thirdCar := new(car)
  thirdCar.model = "Seat Leon"
  fmt.Println(thirdCar.model)
  println(thirdCar) // memory address
  fmt.Println(thirdCar)
}
