package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Degree string
}

func (p Person) PersonInfo() string {
	return fmt.Sprintf("Имя: %s\nВозраст: %d лет\nСпециальность: %s\n\n", p.Name, p.Age, p.Degree)
}

func FillPeopleInTheRoom() []Person {
	people := []Person{
		{"Катя", 20, "Управление в инж. бизнесе"},
		{"Динора", 22, "Врач"},
		{"Настя", 18, "Дизайнер"},
	}
	return people
}

func PrintPeopleInTheRoom() {
	people := FillPeopleInTheRoom()
	fmt.Print("Проживающие в комнате:\n\n")
	for i := 0; i < len(people); i++ {
		fmt.Printf("Проживающий %d:\n\n", i+1)
		fmt.Print(people[i].PersonInfo())
	}
}
