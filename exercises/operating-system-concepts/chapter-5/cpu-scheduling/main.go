package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	cpuscheduling "github.com/karlosdaniel451/cpu-scheduling/cpu_scheduling"
	"github.com/karlosdaniel451/cpu-scheduling/reader"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("invalid input: insuficient parameters: pass the filepath of the" +
			" tasks file and the scheduling policy (fcfs, sjf, pr, rr or pwrr)")
	}

	tasksFilePath := os.Args[1]
	schedulingPolicy := os.Args[2]

	tasks, err := reader.ReadTasksFromCSV(tasksFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var taskManager cpuscheduling.TaskManager

	switch schedulingPolicy {
	case "fcfs":
		taskManager = cpuscheduling.NewTaskManagerFirstComeFirstServed()
		fmt.Println("\nScheduling policy to be used: First-Come First Served.")
	case "sjf":
		taskManager = cpuscheduling.NewTaskManagerShortestJobFirst()
		fmt.Println("\nScheduling policy to be used: Shortest Job first.")
	default:
		log.Fatal("invalid input: unknown scheduling policy")
	}

	for _, task := range tasks {
		taskManager.CreateTask(task.Name, task.Priority, task.CpuBurst)
	}

	fmt.Println("\ntasks to be executed:")
	tasksToBeExecuted := taskManager.GetRunnableTasks()
	sort.Slice(tasksToBeExecuted, func(i, j int) bool {
		return tasksToBeExecuted[i].Tid < tasksToBeExecuted[j].Tid
	})
	for _, task := range tasksToBeExecuted {
		fmt.Printf("%+v\n", *task)
	}

	fmt.Println()
	taskManager.Schedule()
}
