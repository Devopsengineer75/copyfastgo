package main

import (
	"fmt"

	"github.com/Devopsengineer75/copyfastgo/internal/lib"
)

func main() {

	//var name string //nil
	//var name2 string = "Florent"
	//var name3 = "Florent"
	name4 := "Florent"
	age := 30

	fmt.Println("Hello " + name4 + " j'ai " + lib.ConvertIntToString(age) + " ans")

	fmt.Println(lib.Division(3, 0))
}
