package main

import (
	"context"
	"fmt"
	pgx "github.com/jackc/pgx/v5"
	"os"
)

type Person struct {
	Name    string
	Surname string
	Room    string
	Age     int
}

func (p Person) PersonInfo() string {
	return fmt.Sprintf("Имя: %s\nФамилия: %s\nКомната: %s\nВозраст: %d лет\n\n", p.Name, p.Surname, p.Room, p.Age)
}

func GetPeopleFromDB() ([]Person, error) {
	urlExample := "postgres://Home:123@localhost:5436/test_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT name, surname, room, age FROM people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.Name, &person.Surname, &person.Room, &person.Age); err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return people, nil
}

func PrintPeopleInTheRoom() {
	people, err := GetPeopleFromDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting people from the database: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("\033[1mПроживающие в комнате:\033[0m\n\n")
	for i := 0; i < len(people); i++ {
		fmt.Printf("Проживающий %d:\n\n", i+1)
		fmt.Print(people[i].PersonInfo())
	}
}
