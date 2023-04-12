-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2023-04-12 06:01:49.873

-- tables
-- Table: Daily
CREATE TABLE Daily (
    DailyID char(5)  NOT NULL,
    Date date  NULL,
    S_Start timestamp  NULL,
    S_EndTime timestamp  NULL,
    SleepQuality int  NULL,
    W_Amount int  NULL,
    User_UserID char(5)  NOT NULL,
    CONSTRAINT Daily_pk PRIMARY KEY (DailyID)
);

-- Table: Habit
CREATE TABLE Habit (
    HabitID char(5)  NOT NULL,
    HabitName varchar(30)  NOT NULL,
    HabitDescription varchar(100)  NOT NULL,
    StartDate date  NOT NULL,
    EndDate date  NOT NULL,
    ReminderTime time  NOT NULL,
    ReminderFrequency int  NOT NULL,
    User_UserID char(5)  NOT NULL,
    CONSTRAINT Habit_pk PRIMARY KEY (HabitID)
);

-- Table: ToDo
CREATE TABLE ToDo (
    TaskID char(5)  NOT NULL,
    TaskName varchar(30)  NOT NULL,
    DueDate datetime  NOT NULL,
    PriorityLevel int  NOT NULL,
    User_UserID char(5)  NOT NULL,
    CONSTRAINT ToDo_pk PRIMARY KEY (TaskID)
);

-- Table: Tracking
CREATE TABLE Tracking (
    TrackingID char(5)  NOT NULL,
    TrackingDate date  NOT NULL,
    CompletionStatus boolean  NOT NULL,
    Habit_HabitID char(5)  NOT NULL,
    CONSTRAINT Tracking_pk PRIMARY KEY (TrackingID)
);

-- Table: User
CREATE TABLE User (
    UserID char(5)  NOT NULL,
    Name varchar(30)  NOT NULL,
    Email varchar(50)  NOT NULL,
    Password varchar(30)  NOT NULL,
    CONSTRAINT User_pk PRIMARY KEY (UserID)
);

-- foreign keys
-- Reference: Daily_User (table: Daily)
ALTER TABLE Daily ADD CONSTRAINT Daily_User FOREIGN KEY Daily_User (User_UserID)
    REFERENCES User (UserID);

-- Reference: Habit_User (table: Habit)
ALTER TABLE Habit ADD CONSTRAINT Habit_User FOREIGN KEY Habit_User (User_UserID)
    REFERENCES User (UserID);

-- Reference: ToDo_User (table: ToDo)
ALTER TABLE ToDo ADD CONSTRAINT ToDo_User FOREIGN KEY ToDo_User (User_UserID)
    REFERENCES User (UserID);

-- Reference: Tracking_Habit (table: Tracking)
ALTER TABLE Tracking ADD CONSTRAINT Tracking_Habit FOREIGN KEY Tracking_Habit (Habit_HabitID)
    REFERENCES Habit (HabitID);

-- End of file.

