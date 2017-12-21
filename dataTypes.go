package main

func main(){
  var myNum int;
  myNum = 1;
  println("My num is", myNum)

  var myFloat float32 = 31.
  println("My float is ",myFloat)

  myComplex := complex(2,3)
  println(myComplex)
  println(real(myComplex), imag(myComplex))
}
