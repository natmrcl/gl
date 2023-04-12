// Contains the type definitions used in GO, all of the types are an
// implementation of the value object concept in DDD
package CRUDtypes

import (
	"time"
)

type PriorityLevel struct {
	value int
}

func CreatePriorityLevel(level int) *PriorityLevel {
	if level < 0 || level > 5 {
		return nil
	}
	pl := PriorityLevel{value: level}
	return &pl
}

type UserID struct {
	Id string
}

type DueDate struct {
	day time.Time
}

func CreateDueDate(year int, month int, day int, hour int, min int) *DueDate {
	// TODO: CHECK FOR CONSTRAINTS
	if year < 2000 || month < 0 || month > 12 || day < 0 || day > 31 || hour < 0 || hour > 24 || min < 0 || min > 60 {
		return nil
	}
	currentTime := time.Now()
	dueDateTime := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.UTC)
	if dueDateTime.Before(currentTime) {
		return nil
	}
	dueDateObj := DueDate{day: dueDateTime}
	return &dueDateObj
}

type TaskStruct struct {
	TaskID        UserID
	TaskName      string
	DueDate       DueDate
	PriorityLevel PriorityLevel
	UserID        UserID
}
