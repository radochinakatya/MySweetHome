package main

import "fmt"

type Appliances struct {
	Name  string
	Count int
}

func (p Appliances) AppliancesInfo() string {
	return fmt.Sprintf("Название: %s\nКоличество: %d\n\n", p.Name, p.Count)
}

func FillAppliancesInStock() []Appliances {
	appliances := []Appliances{
		{"Холодильник", 1},
		{"Чайник", 1},
		{"Вентилятор", 1},
		{"Удлинитель", 4},
		{"Утюг", 1},
	}
	return appliances
}

func PrintAppliancesInStock() {
	appliances := FillAppliancesInStock()
	fmt.Print("Бытовая техника и др:\n")
	for i := 0; i < len(appliances); i++ {
		fmt.Print(appliances[i].AppliancesInfo())
	}
}
