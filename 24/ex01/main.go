package main

import "fmt"

func main() {
	type item struct {
		id    int
		name  string
		price int
	}
	type game struct {
		item
		genre string
	}
	games := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}
	for _, v := range games {
		fmt.Printf("%-15d %-15q %-20s $%d", v.id, v.item.name, v.name, v.id)
		fmt.Println()
	}
	println(games)
}
