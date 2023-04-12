package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

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

	fmt.Println("Please Select a function:\n1.Create a new task\n2.Create a new habit.")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if strings.Compare(str, "1") == 0 {
		// Function new task
	} else if strings.Compare(str, "2") == 0 {
		// Function new habit
	}

}
