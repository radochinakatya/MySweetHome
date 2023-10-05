package main

import "fmt"

type Item struct {
	Name  string
	Size  string
	Count int
}

func (p Item) ItemInfo() string {
	return fmt.Sprintf("Название: %s\nРазмеры: %s лет\nКоличество: %d\n\n", p.Name, p.Size, p.Count)
}
func FillFurnitureInStock() []Item {
	furniture := []Item{
		{"Кровать", "2 x 0.5 метров", 3},
		{"Стул", "0.3 x 0.3 метров", 3},
		{"Стол", "2 x 0.5 метров", 2},
		{"Мусорка", "0.1 x 0.1 метров", 1},
		{"Полка", "0.2 x 0.3 метров", 5},
	}
	return furniture
}

func PrintFurnitureInStock() {
	furniture := FillFurnitureInStock()
	fmt.Print("Мебель в наличии:\n\n")
	for i := 0; i < len(furniture); i++ {
		fmt.Print(furniture[i].ItemInfo())
	}
}
