package main

import (
	"database/CRUDtypes"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Habit struct {
	HabitID           string
	HabitName         string
	HabitDescription  string
	StartDate         time.Time
	EndDate           time.Time
	ReminderTime      time.Time
	ReminderFrequency int
	UserID            string
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

func CreateHabit(db *sql.DB, habit Habit) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO habit(HabitName, HabitDescription, StartDate, EndDate, ReminderTime, ReminderFrequency, User_UserID) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(habit.HabitID, habit.HabitName, habit.HabitDescription, habit.StartDate, habit.EndDate, habit.ReminderTime, habit.ReminderFrequency, habit.UserID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
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
	// reader := bufio.NewReader(os.Stdin)
	// str, prefix, err := reader.ReadLine()
	// fmt.Println(str, prefix)
	var str string
	fmt.Scanln(&str)

	choice, err := strconv.Atoi(str)
	if choice == 1 {
		// Function new task

	} else if choice == 2 {
		// Function new habit
		// contoh test
		// habit := Habit{
		// 	HabitID:           "H0001",
		// 	HabitName:         "Reading",
		// 	HabitDescription:  "Do some reading every night",
		// 	StartDate:         time.Now(),
		// 	EndDate:           time.Now().AddDate(0, 0, 7),
		// 	ReminderTime:      time.Now().Add(time.Hour),
		// 	ReminderFrequency: 2,
		// 	UserID:            "U0001",
		// }

		// rowsAffected, err := CreateHabit(db, habit)
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println("Data Succesfully Add")
		// 	fmt.Printf("Rows affected: %d\n", rowsAffected)
		// }
	}

}
