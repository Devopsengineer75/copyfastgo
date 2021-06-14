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

	chien := zoo.MakeChien("Olaf", "Woaf !")
	chat := zoo.MakeChien("Gros Minet", "Miaou !")
	oiseau := zoo.MakeChien("titi", "Piou !")
	//fmt.Println(chien, chien.Pattes)
	//chien.Say()

	adopte(chien)
	adopte(chat)
	adopte(oiseau)
}

func adopte(animal zoo.IsAnimal) {
	switch a := animal.(type) {
	case *zoo.Chat:
		fmt.Println("Super je peux adopter", a.Name)
	case *zoo.Chien:
		fmt.Println("Super je peux adopter", a.Name)
	case *zoo.Oiseau:
		fmt.Println("Super je peux adopter", a.Name)
	}

}
