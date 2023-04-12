package main

import (
	"bufio"
	"database/CRUDtypes"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Habits struct {
	HabitID string `json:"HabitID"`
}

func newTask(taskName string, PriorityLevel int8, userID string, year int, month int, day int, hour int, min int, db *sql.DB) {
	time.Now() // temp biar compiler ga ngedelete time.
	piorityLevelObj := CRUDtypes.CreatePriorityLevel(int(PriorityLevel))
	TimeOfTask := CRUDtypes.CreateDueDate(year, month, day, hour, min)
	taskID := CRUDtypes.UserID{
		Id: "1",
	}

	userIDObj := CRUDtypes.UserID{
		Id: userID,
	}

	var task = CRUDtypes.TaskStruct{
		TaskID:        taskID,
		PriorityLevel: *piorityLevelObj,
		TaskName:      taskName,
		DueDate:       *TimeOfTask,
		UserID:        userIDObj,
	}
	results, err := db.Query("INSERT INTO todo (TaskID, TaskName, DueDate, PriorityLevel, User_UserID)") // TODO KELARIN INI
	if err != nil {
		fmt.Println("Error Entering Query:", err)
		return
	}
	fmt.Println(results, task)
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
