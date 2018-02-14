package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	thearray := make([]string, 2)
	json.Unmarshal([]byte(`["banco1", "banco2"]`), &thearray)
	fmt.Println(thearray)
	slice := append(thearray[:], "hola")
	fmt.Println(slice)

	stored, err := json.Marshal(slice)
	if err != nil {
		fmt.Print("")
	}

	thearray2 := make([]string, 3)
	json.Unmarshal(stored, &thearray2)
	slice2 := append(thearray2[:], "adios")
	i := 0
	for i < len(slice2) {
		fmt.Println(slice2[i])
		i++
	}
}
