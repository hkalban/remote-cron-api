package model

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Hostname string  `json:"hostname"`
	IP       string  `json:"ip"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Tasks    []*Task `gorm:"many2many:task_servers;"`
}

type Task struct {
	gorm.Model
	Name            string    `json:"name"`
	Command         string    `json:"command"`
	IntervalSeconds int       `json:"interval_seconds"`
	TimeoutSeconds  int       `json:"timeout_seconds"`
	Servers         []*Server `gorm:"many2many:task_servers;"`
	TaskExecutions  []TaskExecution
}

type TaskExecution struct {
	gorm.Model
	TaskID    uint
	Execution []Execution
}

type Execution struct {
	gorm.Model
	TaskExecutionID uint
	ServerID        uint
	Server          Server
	EndedAt         string
	ExitCode        int
	Output          string `gorm:"type:bytea"`
}
