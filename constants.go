package main

const (
  message = "This is a constant, can not be changed"
)

var msg string // Is not constant // var name type

func main(){
  msg = "I changed the value"

  println(message)
  println(msg)
}
