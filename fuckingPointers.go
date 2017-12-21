package main

import "fmt"

func printMsg(message *string){
      fmt.Println("Memory address: ",&message)  //Prints memory address
      fmt.Println("Message:", *message)
      fmt.Println("Memory address of the function variable:", message)

      //Change value of the address
      *message = "I changed!!"
}

func main(){
  msg := "Be water my friend"

  printMsg(&msg) // Returns be water my friend
  fmt.Println(msg) // returns I changed!
}
