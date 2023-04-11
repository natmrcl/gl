package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Habits struct {
	HabitID string `json:"HabitID"`
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/habimate")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	fmt.Println("Database connection successful")

	results, err := db.Query("SELECT HabitID FROM habit")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var habit Habits
		err = results.Scan(&habit.HabitID)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(habit.HabitID)
	}
}
