package zoo

import "fmt"

type Chien struct {
	Pattes int
	word   string
}

func (c *Chien) Say() {
	fmt.Println(c.word)
}

func MakeChien(word string) *Chien {

	return &Chien{
		Pattes: 4,
		word:   word,
	}
}
