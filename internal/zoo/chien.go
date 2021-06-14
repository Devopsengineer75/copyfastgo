package zoo

import "fmt"

type Chien struct {
	pattes int
	word   string
}

func (c *Chien) Say() {
	fmt.Println(c.word)
}

func MakeChien(word string) Chien {

	return &Chien{
		pattes: 4,
		word:   word,
	}
}
