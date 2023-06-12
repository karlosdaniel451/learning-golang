package cpuscheduling

import (
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

func (taskManager *TaskManagerFirstComeFirstServed) GetCpu() *CPU {
	return taskManager.cpu
}

func (taskManager *TaskManagerFirstComeFirstServed) Schedule() {
	schedule(taskManager)
}

func (taskManager *TaskManagerFirstComeFirstServed) pickNextTask() (task *Task, ok bool) {
	// log.Printf("number of runnable tasks: %d", taskManager.runnableTasks.Size())
	return taskManager.runnableTasks.Dequeue()
}
