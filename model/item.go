package model

import "fmt"

type Item struct {

}

func (i *Item) Shout() {
	fmt.Println("Aaaaaa!")
}
