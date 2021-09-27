package scheduler

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/hkalban/remote-cron-api/model"
)

func (s *BaseScheduler) Start() {
	s.loadTasksExistingTasks()

	s.scheduler.StartAsync()
}

// Loads and schedules existing tasks
func (s *BaseScheduler) loadTasksExistingTasks() {
	// Get Tasks from Database
	var tasks []model.Task

	tasks, err := s.services.TaskService.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	// Schedule and execute existing tasks
	if len(tasks) > 0 {
		for _, task := range tasks {
			s.executeTask(task)
		}
	}
}

// Add a new task to the scheduler on demand
func (s *BaseScheduler) AddNewTask(task model.Task) {
	s.executeTask(task)
}

// Schedule and execute a task
func (s *BaseScheduler) executeTask(task model.Task) {
	s.scheduler.Every(task.IntervalSeconds).Seconds().Do(func() {
		taskExecutionID := s.insertTaskExecution(task)
		for _, server := range task.Servers {
			s.taskExecution(task, *server, taskExecutionID)
		}
	})
}

// Execute task's command and update database
func (s *BaseScheduler) taskExecution(task model.Task, server model.Server, taskExecutionID uint) {
	executionEntity := s.insertExecution(taskExecutionID, server)

	// debug logs
	log.Println(
		fmt.Sprintf(
			"Executing task=%s, server=%s command=%s, every=%d seconds",
			task.Name, server.Hostname, task.Command, task.IntervalSeconds))

	args := strings.Fields(task.Command)

	// TODO: execute command remotely instead of on host
	// Prepending ssh <remote-server> <command>
	// args = append([]string{"ssh", server.Hostname}, args...)

	// TODO: execute timeout command
	// if task.TimeoutSeconds > 0 {
	// 	// Prepending "timeout N" to command
	// 	args = append([]string{"timeout", strconv.Itoa(task.TimeoutSeconds)}, args...)
	// }

	log.Println(args)

	// execute command
	cmd := exec.Command(args[0], args[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		// Update repo with execution output and status
		s.updateExecution(executionEntity, string(err.Error()), -1)
		log.Println(err.Error())
		return
	}

	// Log the output for debugging
	log.Println(string(stdout))

	// Update repo with execution output and status
	s.updateExecution(executionEntity, string(stdout), 0)
}

// Insert Task Execution entry into repository
func (s *BaseScheduler) insertTaskExecution(task model.Task) uint {
	taskExecutionEntity := model.TaskExecution{
		TaskID: task.ID,
	}

	s.services.TaskExecutionService.Create(&taskExecutionEntity)

	return taskExecutionEntity.ID
}

// Insert Execution entry into repository
func (s *BaseScheduler) insertExecution(taskExecutionID uint, server model.Server) model.Execution {
	executionEntity := model.Execution{
		TaskExecutionID: taskExecutionID,
		ServerID:        server.ID,
	}

	s.services.ExecutionService.Create(&executionEntity)

	return executionEntity
}

// Update Execution entry in repository
func (s *BaseScheduler) updateExecution(executionEntity model.Execution, output string, exitCode int) error {
	executionEntity.Output = output
	executionEntity.ExitCode = exitCode

	_, err := s.services.ExecutionService.Update(&executionEntity)
	if err != nil {
		return err
	}

	return nil
}
