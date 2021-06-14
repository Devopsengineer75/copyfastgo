package zoo

import "fmt"

type Oiseau struct {
	Pattes int
	word   string
	Name   string
}

func (c *Oiseau) Say() {
	fmt.Println(c.word)
}

func MakeOiseau(Name string, word string) *Oiseau {

	return &Oiseau{
		Name:   Name,
		Pattes: 2,
		word:   word,
	}
}
