package main

import (
	"context"
	"fmt"
	pgx "github.com/jackc/pgx/v5"
	"os"
)

type Appliances struct {
	Name  string
	Count int
}

func (p Appliances) AppliancesInfo() string {
	return fmt.Sprintf("Название: %s\nКоличество: %d\n\n", p.Name, p.Count)
}

func GetAppliancesFromDB() ([]Appliances, error) {
	urlExample := "postgres://Home:123@localhost:5436/test_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT type, count FROM appliances")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appliances []Appliances
	for rows.Next() {
		var appliancesItem Appliances
		if err := rows.Scan(&appliancesItem.Name, &appliancesItem.Count); err != nil {
			return nil, err
		}
		appliances = append(appliances, appliancesItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appliances, nil
}

func PrintAppliancesInStock() {
	appliances, err := GetAppliancesFromDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting people from the database: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("\033[1mБытовая техника и др:\033[0m\n\n")
	for i := 0; i < len(appliances); i++ {
		fmt.Print(appliances[i].AppliancesInfo())
	}
}
