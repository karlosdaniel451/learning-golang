package main

import (
	"fmt"
	"log"
	"os"

	"github.com/karlosdaniel451/cpu-scheduling/cpu_scheduling"
	"github.com/karlosdaniel451/cpu-scheduling/reader"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("invalid input: insuficient parameters: pass the filepath of the" +
			" tasks file")
	}

	tasksFilePath := os.Args[1]

	tasks, err := reader.ReadTasksFromCSV(tasksFilePath)
	if err != nil {
		log.Fatal(err)
	}

	taskManagerFCFS := cpuscheduling.NewTaskManagerFirstComeFirstServed()

	for _, task := range tasks {
		taskManagerFCFS.CreateTask(task.Name, task.Priority, task.CpuBurst)
	}

	fmt.Println("\ntasks to be executed:")
	tasksToBeExecuted := taskManagerFCFS.GetRunnableTasks()
	for _, task := range tasksToBeExecuted {
		fmt.Printf("%+v\n", *task)
	}

	fmt.Println("\nscheduling tasks with First Come First Served policy...")
	taskManagerFCFS.Schedule()
}
