package main

import (
	"fmt"

	"github.com/Devopsengineer75/copyfastgo/internal/lib"
	"github.com/Devopsengineer75/copyfastgo/internal/zoo"
)

func main() {

	//var name string //nil
	//var name2 string = "Florent"
	//var name3 = "Florent"
	name4 := "Florent"
	age := 30

	fmt.Println("Hello " + name4 + " j'ai " + lib.ConvertIntToString(age) + " ans")

	fmt.Println(lib.Division(3, 0))
	fmt.Println(lib.Division(3, 2))

	chien := zoo.MakeChien("Woaf !")
	fmt.Println(chien, chien.Pattes)
	chien.Say()
}
