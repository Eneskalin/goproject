package main

import (
	"app/items"
	"fmt"
)

func main() {
	list := items.ItemList()

	for _, i := range list {
		fmt.Println(i.Name)
	}
}
