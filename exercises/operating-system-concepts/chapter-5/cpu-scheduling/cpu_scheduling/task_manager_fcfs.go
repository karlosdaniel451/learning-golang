package cpuscheduling

import (
	"log"
	"time"

	"github.com/ugurcsen/gods-generic/queues/linkedlistqueue"
)

type TaskManagerFirstComeFirstServed struct {
	cpu           *CPU
	runnableTasks linkedlistqueue.Queue[*Task]
	nextTaskID    int
}

func NewTaskManagerFirstComeFirstServed() *TaskManagerFirstComeFirstServed {
	return &TaskManagerFirstComeFirstServed{
		nextTaskID:    1,
		runnableTasks: *linkedlistqueue.New[*Task](),
		cpu:           newCPU(),
		// scheduler:     RoundRobinScheduler{},
	}
}

func (taskManager *TaskManagerFirstComeFirstServed) GetRunnableTasks() []*Task {
	return taskManager.runnableTasks.Values()
}

func (taskManager *TaskManagerFirstComeFirstServed) GetNextTaskID() int {
	defer func() {
		taskManager.nextTaskID++
	}()

	return taskManager.nextTaskID
}

func (taskManager *TaskManagerFirstComeFirstServed) CreateTask(
	name string,
	priority int,
	cpuBurst int,
) *Task {
	task := &Task{
		Name:              name,
		Tid:               taskManager.GetNextTaskID(),
		Priority:          priority,
		CpuBurst:          cpuBurst,
		RemainingCpuBurst: cpuBurst,
		State:             Runnable,
	}

	taskManager.runnableTasks.Enqueue(task)

	return task
}

func (taskManager *TaskManagerFirstComeFirstServed) Schedule() {
	for {
		nextTaskToBeRun, ok := taskManager.pickNextTask()
		if !ok {
			break
		}
		// FirstComeFirstServed scheduling policy is non-preemptive, so it
		// can only select another process/thread/task to run when there is
		// no one currently running in the CPU
		for {
			isCpuIdle := taskManager.cpu.isIdle()
			if isCpuIdle {
				break
			}
			time.Sleep(time.Second * 1)
			// log.Println("cpu is still not idle")
		}
		previouslyExecutingTask := taskManager.cpu.currentlyExecutingTask
		if previouslyExecutingTask != nil {
			previouslyExecutingTask.State = Terminated
		}
		nextTaskToBeRun.State = Running
		taskManager.cpu.runTask(nextTaskToBeRun, nextTaskToBeRun.CpuBurst)
	}
	// Wait for the last task finish
	for {
		isCpuIdle := taskManager.cpu.isIdle()
		if isCpuIdle {
			break
		}
		time.Sleep(time.Second * 1)
		// log.Println("cpu is still not idle")
	}
}

func (taskManager *TaskManagerFirstComeFirstServed) pickNextTask() (task *Task, ok bool) {
	log.Printf("number of runnable tasks: %d", taskManager.runnableTasks.Size())
	return taskManager.runnableTasks.Dequeue()
}
