package main

import (
	"context"
	"fmt"
	pgx "github.com/jackc/pgx/v5"
	"os"
)

type Item struct {
	Name  string
	Size  string
	Count int
}

func (p Item) ItemInfo() string {
	return fmt.Sprintf("Название: %s\nРазмеры: %s лет\nКоличество: %d\n\n", p.Name, p.Size, p.Count)
}

func GetFurnitureFromDB() ([]Item, error) {
	urlExample := "postgres://Home:123@localhost:5436/test_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT type, size, count FROM furniture")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var furniture []Item
	for rows.Next() {
		var furnitureItem Item
		if err := rows.Scan(&furnitureItem.Name, &furnitureItem.Size, &furnitureItem.Count); err != nil {
			return nil, err
		}
		furniture = append(furniture, furnitureItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return furniture, nil
}

func PrintFurnitureInStock() {
	furniture, err := GetFurnitureFromDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting people from the database: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("\033[1mМебель в наличии:\033[0m\n\n")
	for i := 0; i < len(furniture); i++ {
		fmt.Print(furniture[i].ItemInfo())
	}
}
